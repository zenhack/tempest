package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	configureFlags = flag.NewFlagSet("configure", flag.ExitOnError)

	user       = configureFlags.String("user", "granular", "the user to run as")
	group      = configureFlags.String("group", "granular", "the group to run as")
	prefix     = configureFlags.String("prefix", "/usr/local", "install prefix")
	libexecdir = configureFlags.String("libexecdir", "",
		`path for helper commands (default "${PREFIX}/libexec")`)
	localstatedir = configureFlags.String("localstatedir", "",
		`path to store run-time data (default "${PREFIX}/var/lib")`)
)

func chkfatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func spawnErr(fn func() error) <-chan error {
	ret := make(chan error)
	go func() {
		ret <- fn()
	}()
	return ret
}

func withMyOuts(cmd *exec.Cmd) *exec.Cmd {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func runInDir(dir, bin string, args ...string) error {
	cmd := exec.Command(bin, args...)
	cmd.Dir = dir
	return withMyOuts(cmd).Run()
}

func buildC() error {
	return runInDir("c", "make")
}

func buildElm() error {
	return runInDir("elm", "elm-make", "Main.elm", "--yes", "--output", "../static/index.html")
}

func buildGo() error {
	return runInDir(".", "go", "build", "-v", "-i", "-o", "bin/server", "./go/cmd/server")
}

func cleanC() error {
	return runInDir("c", "make", "clean")
}

func cleanElm() error {
	return runInDir("elm", "rm", "-rf", "elm-stuff", "../static/index.html")
}

func cleanGo() error {
	return runInDir(".", "rm", "-f", "bin/server")
}

func nukeC() error {
	return runInDir(".", "rm", "-f", "c/config.h")
}

func nukeGo() error {
	return runInDir(".", "rm", "-f", "go/internal/config/config.go")
}

// Run configure if its outputs aren't already present.
func maybeConfigure() {
	_, errC := os.Stat("./c/config.h")
	_, errGo := os.Stat("./go/internal/config/config.go")
	if errC == nil && errGo == nil {
		// Config is already present; we're done.
		return
	}
	log.Println("'configure' has not been run; running with default options.")
	run("configure")
}

func runJobs(jobs ...func() error) {
	chans := make([]<-chan error, len(jobs))
	errs := make([]error, len(jobs))
	for i := range jobs {
		chans[i] = spawnErr(jobs[i])
	}
	for i := range chans {
		errs[i] = <-chans[i]
	}
	for i := range errs {
		chkfatal(errs[i])
	}
}

func run(args ...string) {
	switch args[0] {
	case "build":
		maybeConfigure()
		runJobs(buildC, buildElm, buildGo)
	case "run":
		run("build")
		fmt.Fprintln(os.Stderr, "Starting server...")
		chkfatal(withMyOuts(exec.Command("./bin/server")).Run())
	case "clean":
		maybeConfigure()
		runJobs(cleanC, cleanElm, cleanGo)
	case "nuke":
		run("clean")
		runJobs(nukeC, nukeGo)
	case "configure":
		configureFlags.Parse(args[1:])
		if *libexecdir == "" {
			*libexecdir = *prefix + "/libexec"
		}
		if *localstatedir == "" {
			*localstatedir = *prefix + "/var/lib"
		}
		file, err := os.Create("./go/internal/config/config.go")
		chkfatal(err)
		defer file.Close()
		_, err = fmt.Fprintf(file,
			`
package conifg

const (
	User = %q
	Group = %q
	Prefix = %q
	Libexecdir = %q
	Localstatedir = %q
)
`,
			*user,
			*group,
			*prefix,
			*libexecdir,
			*localstatedir,
		)
		chkfatal(err)

		file, err = os.Create("./c/config.h")
		chkfatal(err)
		defer file.Close()
		_, err = fmt.Fprintf(file,
			`
#ifndef GRANULAR_CONFIG_H
#define GRANULAR_CONFIG_H

#define PREFIX %q
#define LIBEXECDIR %q
#define LOCALSTATEDIR %q

#endif
`,
			*prefix,
			*libexecdir,
			*localstatedir,
		)
		chkfatal(err)
	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", args[0])
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		run("build")
	} else {
		run(os.Args[1:]...)
	}
}
