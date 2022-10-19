/* This program is responsible for actually setting up the grain sandbox.
 * Sandstorm invokes it as:
 *
 * sandstorm-sandbox-launcher <package-id> <grain-id>
 *
 * It configures a sandbox using the directories for that package & grain, and
 * then inovkes execveat() to start the `sandstorm-sandbox-agent` executable, which then
 * takes over starting up and interfacing with the grain proper.
 *
 * This program is written in C, rather than Go, because:
 *
 * 1. There have historically been many bugs and gotchas around multi-threaded
 *    programs that drop elevated privileges, and there is no such thing as a
 *    single-threaded Go program.
 * 2. More broadly, this is an area of the code where we aren't really processing
 *    user input, and we are more concerned with polluting the sandbox with ambient
 *    state from the operating system -- so it is better to not have a complex language
 *    runtime underneath us.
 */
#define _GNU_SOURCE

#include "config.h"

/* bool */
#include <stdbool.h>

#include <sys/syscall.h>

/* unshare */
#include <sched.h>

/* misc networking stuff: */
#include <arpa/inet.h>
#include <sys/socket.h>
#include <linux/sockios.h>
#include <net/if.h>

/* errno */
#include <errno.h>

/* sigprocmask, sigfillset */
#include <signal.h>

/* mount */
#include <sys/mount.h>

/* makedev (for use with mknod) */
#include <sys/sysmacros.h>

/* exit */
#include <stdlib.h>

/* fprintf, stderr, perror */
#include <stdio.h>

/* strlen */
#include <string.h>

/* isxdigit */
#include <ctype.h>

/* mknod, open, fork, waitpid, fexecve */
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/stat.h>
#include <sys/syscall.h>
#include <fcntl.h>
#include <unistd.h>

/* prctl and associated constants */
#include <sys/prctl.h>

#define SANDSTORM_STATE   LOCALSTATEDIR "/sandstorm"
#define SANDSTORM_LIBEXEC LIBEXECDIR    "/sandstorm"

#define IMAGE_DIR     SANDSTORM_STATE   "/apps"
#define SANDBOX_DIR   SANDSTORM_STATE   "/grains"
#define AGENT_PATH    SANDSTORM_LIBEXEC "/sandstorm-sandbox-agent"
#define CHROOT_MNT    SANDSTORM_STATE   "/mnt"

#define PKG_ID_SIZE 32
#define GRAIN_ID_SIZE 22

#define REQUIRE(condition) \
	if (!(condition)) do { \
		panic(__FILE__, __LINE__, #condition); \
	} while(0)

void panic(const char *file, int line, const char *check) {
	if(errno != 0) {
		perror("ERRNO: ");
	}
	fprintf(stderr, "FATAL: %s line %d: REQUIRE(%s) failed\n", file, line, check);
	exit(1);
}

void require_valid_pkg_id(const char *str) {
	REQUIRE(strlen(str) == PKG_ID_SIZE);
	while(*str) {
		REQUIRE(isxdigit(*str));
		str++;
	}
}

void require_valid_grain_id(const char *str) {
	REQUIRE(strlen(str) == GRAIN_ID_SIZE);
	while(*str) {
		REQUIRE(isalpha(*str) || isdigit(*str));
		str++;
	}
}

int main(int argc, char **argv) {
	REQUIRE(argc == 3);
	require_valid_pkg_id(argv[1]);
	require_valid_grain_id(argv[2]);

	const char *image_id = argv[1];
	const char *sandbox_id = argv[2];

	/* Get an fd for the agent executable, which we'll execveat() once we're in the sandbox. */
	int agent_fd = open(AGENT_PATH, O_RDONLY);
	REQUIRE(agent_fd >= 0);

	/* Set it to close-on-exec, so we don't leak the fd into the sandbox. */
	REQUIRE(fcntl(agent_fd, F_SETFD, fcntl(agent_fd, F_GETFD) | FD_CLOEXEC) != -1);

	/* Unshare basically all of the namespaces. We leave out user namespaces, since they
	   aren't universally supported. */
	/* TODO: make sure we're cleaning up the resources for each of these:

	   - [ ] files
	   - [ ] fs
	   - [ ] cgroup
	   - [ ] ipc
	   - [ ] net
	   - [ ] pid
	   - [ ] uts
	   - [ ] sysvsem

	   */
	REQUIRE(unshare(
		CLONE_NEWNS |
		CLONE_FILES |
		CLONE_FS |
		CLONE_NEWCGROUP |
		CLONE_NEWIPC |
		CLONE_NEWNET |
		CLONE_NEWPID |
		CLONE_NEWUTS |
		CLONE_SYSVSEM) == 0);

	/* No, really, unshare the mounts. See the "SHARED SUBTREES" section of mount_namespaces(7)
	   for details. */
	REQUIRE(mount("", "/", "", MS_REC|MS_PRIVATE, "") == 0);

	/* Set up a loopback interface in the new network namespace. */
	{
		/* Socket to make ioctls: */
		int sockfd = socket(AF_INET, SOCK_DGRAM, IPPROTO_IP);
		REQUIRE(sockfd >= 0);

		/* Set the address of "lo": */
		struct ifreq ifr;
		memset(&ifr, 0, sizeof ifr);
		strcpy(ifr.ifr_ifrn.ifrn_name, "lo");
		struct sockaddr_in *addr = (struct sockaddr_in *)(&ifr.ifr_ifru.ifru_addr);
		addr->sin_family = AF_INET;
		addr->sin_addr.s_addr = htonl(0x7f000001); /* 127.0.0.1 */
		REQUIRE(ioctl(sockfd, SIOCSIFADDR, &ifr) >= 0);

		/* ...and set flags to enable it: */
		memset(&ifr.ifr_ifru, 0, sizeof ifr.ifr_ifru);
		ifr.ifr_ifru.ifru_flags = IFF_LOOPBACK | IFF_UP | IFF_RUNNING;
		REQUIRE(ioctl(sockfd, SIOCSIFFLAGS, &ifr) >= 0);

		close(sockfd);
	}

	/* Mount the image read only, then mount the sandbox's storage in the image's /var. */
	REQUIRE(chdir(IMAGE_DIR) == 0);
	REQUIRE(mount(image_id, CHROOT_MNT, "", MS_BIND, "") == 0);
	REQUIRE(mount("", CHROOT_MNT, "", MS_REMOUNT|MS_BIND|MS_RDONLY, "") == 0);
	REQUIRE(chdir(SANDBOX_DIR) == 0);
	REQUIRE(chdir(sandbox_id) == 0);
	REQUIRE(mount("sandbox", CHROOT_MNT "/var", "", MS_BIND, "") == 0);

	/* Bind-mount the host's cpuinfo. It's typical for runtime environments to inspect this. */
	REQUIRE(mount("/proc/cpuinfo", CHROOT_MNT "/proc/cpuinfo", "", MS_BIND, "") == 0);

	/* Supply a small /tmp. */
	REQUIRE(mount("none", CHROOT_MNT "/tmp", "tmpfs", MS_NODEV|MS_NOSUID, "size=16m") == 0);

	/* Set up /dev; a read-only tmpfs with a minimal set of devices. */
	REQUIRE(mount("none", CHROOT_MNT "/dev", "tmpfs", MS_NOSUID, "") == 0);
	REQUIRE(mknod(CHROOT_MNT "/dev/null",    S_IFCHR|0666, makedev(1, 3)) == 0);
	REQUIRE(mknod(CHROOT_MNT "/dev/zero",    S_IFCHR|0666, makedev(1, 5)) == 0);
	REQUIRE(mknod(CHROOT_MNT "/dev/random",  S_IFCHR|0666, makedev(1, 8)) == 0);
	REQUIRE(mknod(CHROOT_MNT "/dev/urandom", S_IFCHR|0666, makedev(1, 9)) == 0);
	REQUIRE(mount("", CHROOT_MNT "/dev", "", MS_REMOUNT|MS_RDONLY|MS_NOSUID, "") == 0);

	/* Close all file descriptors, except for:

	   - stdout & stderr -- these are logged by the supervisor.
	   - fd #3, which is the rpc socket
	   - agent_fd, which we still need to pass to execveat when we're done. It's close-on-exec,
	     so this is fine.

	We don't check the errors here; on linux close() can't actually fail to close the file
	descriptor, and none of the other errors are relevant to us. */
	close(0);
	long max_fds = sysconf(_SC_OPEN_MAX);
	for(int i = 4; i < max_fds; i++) {
		if(i != agent_fd) {
			close(i);
		}
	}

	/* Drop all posix caps from the ambient set; we don't want any of this authority to be
	   inherited by the sandbox. */
	REQUIRE(prctl(PR_CAP_AMBIENT, PR_CAP_AMBIENT_CLEAR_ALL, 0, 0, 0) == 0);

	REQUIRE(prctl(PR_SET_NO_NEW_PRIVS, 1, 0, 0, 0) == 0);


	/* Make sure we're inside the sandbox's root. */
	REQUIRE(chdir(CHROOT_MNT) == 0);

	/* We use pivot_root below to swap the grain's directory in for our new root directory.
	 * pivot_root(new, old) swaps the root mount out for the mount at `new`, and
	 * simultaneously mounts the old root at `old`. We provide CHROOT_MNT as both
	 * paths, which has the weird effect of making CHROOT_MNT our new root, but
	 * mounting the old root on top of it.
	 *
	 * I don't really know where else to put it; there's no obvious place it could
	 * go inside the grain's storage. But this has the unfortunate consequence that
	 * we can't just umount / to get rid of it. Instead, before pivoting we grab a
	 * reference to the old root, which we can use for the umount after it becomes
	 * unaddressable.
	 *
	 * N.B. I believe we should be able to make this work by using mount with MS_MOVE
	 * in flags, but I'm having trouble debugging that version and this is what the
	 * old sandstorm did, so let's just stick with it.
	 *
	 * Once we're in, we're in somewhat dangerous territory, since now the app
	 * controls our filesystem's contents. Accordingly, this as near to the end of
	 * the process as it can be. */
	int old_root = open("/", O_RDONLY | O_DIRECTORY | O_CLOEXEC);
	REQUIRE(old_root >= 0);
	REQUIRE(syscall(SYS_pivot_root, CHROOT_MNT, CHROOT_MNT) == 0);
	REQUIRE(fchdir(old_root) == 0);
	REQUIRE(umount2(".", MNT_DETACH) == 0);

	/* Now actually move into the new root and drop the reference: */
	REQUIRE(chdir("/") == 0);
	REQUIRE(close(old_root) == 0);

	/* TODO: install a seccomp-bpf filter. */

	pid_t pid = fork();
	REQUIRE(pid != -1);
	if(pid != 0) {
		/* parent */
		/* Close the remaining file descriptors. */
		close(1);
		close(2);
		close(3);
		close(agent_fd);

		int wstatus;
		/* FIXME: handle failures from waitpid (EINTR mainly). */
		waitpid(pid, &wstatus, 0);
		return WEXITSTATUS(wstatus);
	} else {
		/* child. We're now pid 1 in the new pid namespace. We'll fork again, execing the
		   agent in the child, and reaping processes in the parent. We can't have the
		   agent do this itself, since it's written in go and the runtime does weird
		   things with clone(). */

		/* NOTE: This block is adapted from the code at http://ewontfix.com/14/. Which is
		   copyright (c) Rich Felker, 2014. The following (standard MIT) license applies:

		   Permission is hereby granted, free of charge, to any person obtaining a copy of
		   this software and associated documentation files (the "Software"), to deal in the
		   Software without restriction, including without limitation the rights to use,
		   copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
		   Software, and to permit persons to whom the Software is furnished to do so,
		   subject to the following conditions:

		       The above copyright notice and this permission notice shall be included in
		       all copies or substantial portions of the Software.

		   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
		   IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
		   FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
		   COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN
		   AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
		   WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE. */


		/* First, make sure something like SIGCHLD doesn't kill us: */
		sigset_t set;
		int status;
		REQUIRE(sigfillset(&set) == 0);
		REQUIRE(sigprocmask(SIG_BLOCK, &set, 0) == 0);

		pid = fork();
		REQUIRE(pid >= 0);
		if(pid != 0) {
			/* We're the parent. First, close the remaining file descriptors: */
			close(1);
			close(2);
			close(3);
			close(agent_fd);

			/* We *should* be pid 1 in the new pid namespace: */
			REQUIRE(getpid() == 1);

			/* So start acting like init and reap processes: */
			while(true) {
				pid_t reaped_pid = wait(&status);
				if(reaped_pid == pid) {
					/* The agent exited; stop. */
					/* TODO: think about how to report this. right now
					   we're exiting(2), which may be reasonable, but
					   we should decide and document. */
					exit(2);
				}
			}
		}

		/* We're the child. Reset the signal mask and launch the agent. */
		REQUIRE(sigprocmask(SIG_UNBLOCK, &set, 0) == 0);

		/* TODO: document what's going on here. */
		REQUIRE(setsid() != -1);
		//REQUIRE(setpgid(0, 0) == 0);

		char *const agent_argv[] = {"sandbox-agent", NULL};
		char *const agent_envp[] = {NULL};
		syscall(SYS_execveat, agent_fd, "", agent_argv, agent_envp, AT_EMPTY_PATH);

		/* Exec failed, bail out: */
		perror("execveat");
		return 1;
	}
}

/* vim: set ts=8 sw=8 tw=100 noet : */
