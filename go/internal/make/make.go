package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"time"
)

type Config struct {
	User, Group   string
	Prefix        string
	ExecPrefix    string
	Bindir        string
	Libexecdir    string
	Localstatedir string

	WithGoSandstorm string
	WithGoCapnp     string
}

func getUid(name string) int {
	u, err := user.Lookup(name)
	chkfatal(err)
	id, err := strconv.Atoi(u.Uid)
	chkfatal(err)
	return id
}

func getGid(name string) int {
	g, err := user.LookupGroup(name)
	chkfatal(err)
	id, err := strconv.Atoi(g.Gid)
	chkfatal(err)
	return id
}

func (c *Config) ParseFlags(args []string, name string, errorHandling flag.ErrorHandling) {
	fs := flag.NewFlagSet(name, errorHandling)
	fs.StringVar(&c.User, "user", "sandstorm", "the user to run as")
	fs.StringVar(&c.Group, "group", "sandstorm", "the group to run as")

	fs.StringVar(&c.Prefix, "prefix", "/usr/local", "install prefix")
	fs.StringVar(&c.ExecPrefix, "exec-prefix", "", "executable prefix (default ${PREFIX})")
	fs.StringVar(&c.Bindir, "bindir", "", "path for executables (default ${EXEC_PREFIX}/bin)")
	fs.StringVar(&c.Libexecdir, "libexecdir", "",
		`path for helper commands (default "${PREFIX}/libexec")`)
	fs.StringVar(&c.Localstatedir, "localstatedir", "",
		`path to store run-time data (default "${PREFIX}/var/lib")`)

	fs.StringVar(&c.WithGoSandstorm, "with-go-sandstorm", "", "path to go.sandstorm source")
	fs.StringVar(&c.WithGoCapnp, "with-go-capnp", "", "path to go-capnp source")

	// currently unused, but permitted, for compatibility with gnu coding guidelines/autoconf.
	fs.String("sbindir", "", "unused")
	fs.String("sysconfdir", "", "unused")
	fs.String("sharedstatedir", "", "unused")
	fs.String("runstatedir", "", "unused")
	fs.String("libdir", "", "unused")
	fs.String("includedir", "", "unused")
	fs.String("oldincludedir", "", "unused")
	fs.String("datarootdir", "", "unused")
	fs.String("datadir", "", "unused")
	fs.String("infodir", "", "unused")
	fs.String("mandir", "", "unused")
	fs.String("docdir", "", "unused")
	fs.String("htmldir", "", "unused")
	fs.String("dvidir", "", "unused")
	fs.String("pdfdir", "", "unused")
	fs.String("psdir", "", "unused")

	fs.Parse(args[1:])

	if c.ExecPrefix == "" {
		c.ExecPrefix = c.Prefix
	}
	if c.Bindir == "" {
		c.Bindir = c.ExecPrefix + "/bin"
	}
	if c.Libexecdir == "" {
		c.Libexecdir = c.Prefix + "/libexec"
	}
	if c.Localstatedir == "" {
		c.Localstatedir = c.Prefix + "/var/lib"
	}
}

func (c Config) GoSrc() string {
	return fmt.Sprintf(`package config

const (
	User = %q
	Group = %q
	Prefix = %q
	Libexecdir = %q
	Localstatedir = %q
)
`,
		c.User,
		c.Group,
		c.Prefix,
		c.Libexecdir,
		c.Localstatedir,
	)
}

func (c Config) CSrc() string {
	return fmt.Sprintf(`
#ifndef SANDSTORM_CONFIG_H
#define SANDSTORM_CONFIG_H

#define PREFIX %q
#define LIBEXECDIR %q
#define LOCALSTATEDIR %q

#endif
`,
		c.Prefix,
		c.Libexecdir,
		c.Localstatedir,
	)
}

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

func installExe(cfg Config, exe, dir, caps string) {
	destDir := os.Getenv("DESTDIR")
	src, err := os.Open("./_build/" + exe)
	chkfatal(err)
	defer src.Close()
	dstPathDir := destDir + dir + "/"
	chkfatal(os.MkdirAll(dstPathDir, 0755))
	dstPath := dstPathDir + exe
	dst, err := os.OpenFile(dstPath, os.O_CREATE|os.O_RDWR, 0750)
	chkfatal(err)
	defer dst.Close()
	_, err = io.Copy(dst, src)
	chkfatal(err)
	chkfatal(os.Chown(dstPath, 0, getGid("sandstorm")))
	if caps != "" {
		chkfatal(withMyOuts(exec.Command("setcap", caps, dstPath)).Run())
	}
}

func buildC() error {
	log.Println("Building C executable")
	return runInDir("c", "make")
}

func buildCapnp(r *BuildRecord) {
	log.Println("Compiling capnp schema")
	c := readConfig()
	dirs := []string{
		"capnp",
		"capnp/internal",
	}

	for _, d := range dirs {
		files, err := filepath.Glob(d + "/*.capnp")
		chkfatal(err)

		for _, file := range files {
			dir := file[:len(file)-len(".capnp")]
			err := os.MkdirAll(dir, 0755)
			chkfatal(err)
			cmd := exec.Command("capnp",
				"compile",
				"-o-",
				"--src-prefix="+d+"/",
				"-I", c.WithGoCapnp+"/std",
				"-I", c.WithGoSandstorm+"/capnp",
				file,
			)
			cgr, err := cmd.Output()
			chkfatal(err)
			cgrPath := file + ".cgr"
			oldSig, ok := r.Files[cgrPath]
			if !ok || oldSig.Stamp.Size != int64(len(cgr)) {
				hash := sha256.Sum256(cgr)
				if bytes.Compare(hash[:], oldSig.Hash) != 0 {
					log.Printf("Generating go code for %q", file)
					chkfatal(os.WriteFile(cgrPath, cgr, 0644))
					cmd := exec.Command("capnpc-go")
					cmd.Dir = dir
					cmd.Stdin, err = os.Open(cgrPath)
					chkfatal(err)
					chkfatal(withMyOuts(cmd).Run())
					chkfatal(r.RecordFile(cgrPath))
				}
			}
		}
	}
}

func buildWebui() error {
	// Build the webassembly binary:
	log.Println("Building wasm binary")
	err := runInDir("go/cmd/webui",
		"tinygo", "build",
		"-target", "wasm",
		"-panic", "trap",
		"-no-debug",
		"-o=../../internal/server/embed/webui.wasm")
	if err != nil {
		return err
	}

	// Copy the js shim. FIXME: be smarter about the source location;
	// this will fail if tinygo is installed via a different path.
	// The stock Go toolchain has this at a location relative to $GOROOT,
	// but I don't know how to adaptively find it for tinygo.
	return copyFile(
		"go/internal/server/embed/wasm_exec.js",
		"/usr/lib/tinygo/targets/wasm_exec.js",
	)
}

func copyFile(dest, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}

func buildGo() error {
	r := GetBuildRecord()
	buildCapnp(r)
	r.Save()

	err := buildWebui()
	if err != nil {
		return err
	}
	exes := []struct {
		name   string
		static bool
	}{
		{"sandstorm-legacy-tool", false},
		{"sandstorm-next", false},
		{"sandstorm-sandbox-agent", true},
	}
	for _, exe := range exes {
		err = compileGoExe(exe.name, exe.static)
		if err != nil {
			return err
		}
	}
	return nil
}

func compileGoExe(name string, static bool) error {
	log.Printf("Compiling go executable %q (static = %v)", name, static)
	cmd := exec.Command("go", "build", "-v", "-o", "_build/"+name, "./go/cmd/"+name)
	cmd.Env = append(cmd.Env, os.Environ()...)
	if static {
		cmd.Env = append(cmd.Env, "CGO_ENABLED=0")
	} else {
		cmd.Env = append(cmd.Env, "CGO_ENABLED=1")
	}
	return withMyOuts(cmd).Run()
}

// Run configure if its outputs aren't already present.
func maybeConfigure() {
	_, errC := os.Stat("./c/config.h")
	_, errGo := os.Stat("./go/internal/config/config.go")
	_, errJson := os.Stat("./config.json")
	if errC == nil && errGo == nil && errJson == nil {
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
		chkfatal(os.MkdirAll("_build", 0755))
		runJobs(buildC, buildGo)
	case "run":
		run("build")
		fmt.Fprintln(os.Stderr, "Starting server...")
		chkfatal(withMyOuts(exec.Command("./bin/server")).Run())
	case "configure":
		cfg := &Config{}
		cfg.ParseFlags(args, "configure", flag.ExitOnError)
		chkfatal(ioutil.WriteFile(
			"./go/internal/config/config.go",
			[]byte(cfg.GoSrc()),
			0600))
		chkfatal(ioutil.WriteFile(
			"./c/config.h",
			[]byte(cfg.CSrc()),
			0600))
		jsonData, err := json.Marshal(cfg)
		chkfatal(err)
		chkfatal(ioutil.WriteFile(
			"./config.json",
			jsonData,
			0600))
	case "install":
		run("build")
		c := readConfig()
		installExe(c, "sandstorm-next", c.Bindir, "cap_net_bind_service+ep")
		installExe(c, "sandstorm-sandbox-launcher", c.Libexecdir+"/sandstorm",
			"cap_sys_admin,cap_net_admin,cap_mknod+ep")
		installExe(c, "sandstorm-sandbox-agent", c.Libexecdir+"/sandstorm", "")
		chkfatal(os.MkdirAll(c.Localstatedir+"/sandstorm/mnt", 0700))
	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", args[0])
		os.Exit(1)
	}
}

func readConfig() Config {
	var c Config
	data, err := ioutil.ReadFile("config.json")
	chkfatal(err)
	chkfatal(json.Unmarshal(data, &c))
	return c
}

func main() {
	if len(os.Args) < 2 {
		run("build")
	} else {
		run(os.Args[1:]...)
	}
}

const buildRecordPath = "_build/build_record.json"

func GetBuildRecord() *BuildRecord {
	empty := &BuildRecord{
		Files: make(map[string]FileSig),
	}
	data, err := os.ReadFile(buildRecordPath)
	if err != nil {
		return empty
	}
	var ret BuildRecord
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return empty
	}
	return &ret
}

type BuildRecord struct {
	Files map[string]FileSig
}

func (r *BuildRecord) Save() {
	data, err := json.Marshal(r)
	chkfatal(err)
	chkfatal(os.WriteFile(buildRecordPath, data, 0644))
}

func (r *BuildRecord) IsModified(path string) bool {
	stamp, err := StampFile(path)
	if err != nil {
		return true
	}
	sig, ok := r.Files[path]
	if !ok {
		return true
	}

	return stamp != sig.Stamp
}

func (r *BuildRecord) RecordFile(path string) error {
	stamp, err := StampFile(path)
	if err != nil {
		return err
	}
	hash, err := HashFile(path)
	if err != nil {
		return err
	}
	r.Files[path] = FileSig{
		Stamp: stamp,
		Hash:  hash,
	}
	return nil
}

func HashFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func StampFile(path string) (FileStamp, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return FileStamp{}, err
	}
	return FileStamp{
		Size:    fi.Size(),
		Mode:    fi.Mode(),
		ModTime: fi.ModTime(),
	}, nil
}

type FileStamp struct {
	Size    int64
	Mode    fs.FileMode
	ModTime time.Time
}

type FileSig struct {
	Stamp FileStamp
	Hash  []byte
}
