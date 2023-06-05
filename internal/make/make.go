package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	User, Group   string
	Prefix        string
	ExecPrefix    string
	Bindir        string
	Libexecdir    string
	Localstatedir string

	WithGoCapnp    string
	WithWasmExecJs string

	TinyGo bool

	// Actual arguments passed to configure:
	Args []string
}

func getGid(name string) int {
	g, err := user.LookupGroup(name)
	chkfatal(err)
	id, err := strconv.Atoi(g.Gid)
	chkfatal(err)
	return id
}

func (c *Config) ParseFlags(args []string, name string, errorHandling flag.ErrorHandling) {
	c.Args = args[1:]

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

	fs.StringVar(&c.WithGoCapnp, "with-go-capnp", "", "path to go-capnp source")
	fs.StringVar(&c.WithWasmExecJs, "with-wasm_exec.js", "", "path to wasm_exec.js")

	fs.BoolVar(&c.TinyGo, "use-tinygo", true, "Use tinygo for webassembly build")

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
#pragma once
#define PREFIX %q
#define LIBEXECDIR %q
#define LOCALSTATEDIR %q
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
	chkfatal(os.Chown(dstPath, 0, getGid(cfg.Group)))
	if caps != "" {
		chkfatal(withMyOuts(exec.Command("setcap", caps, dstPath)).Run())
	}
}

func buildC() error {
	log.Println("Building C executable")
	return runInDir("c", "make")
}

func buildConfig(r *BuildRecord) {
	if r.IsModified("./config.json") {
		cfg := readConfig()
		files := []struct {
			path    string
			content string
		}{
			{
				path:    "./internal/config/config.go",
				content: cfg.GoSrc(),
			},
			{
				path:    "./c/config.h",
				content: cfg.CSrc(),
			},
		}
		for _, f := range files {
			chkfatal(os.WriteFile(f.path, []byte(f.content), 0600))
			r.RecordFile(f.path)
		}
	}
}

func buildCapnp(r *BuildRecord) {
	log.Println("Compiling capnp schema")
	c := readConfig()
	dirs := []string{
		"capnp",
		"internal/capnp",
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
				"-I", "capnp",
				file,
			)
			cmd.Stderr = os.Stderr
			cgr, err := cmd.Output()
			chkfatal(err)
			cgrPath := file + ".cgr"
			oldSig, ok := r.Files[cgrPath]
			if !ok || oldSig.Stamp.Size != int64(len(cgr)) {
				hash := sha256.Sum256(cgr)
				if !bytes.Equal(hash[:], oldSig.Hash) {
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

func findWasmExecJs(cfg Config) (string, error) {
	if cfg.WithWasmExecJs != "" {
		return cfg.WithWasmExecJs, nil
	}
	if cfg.TinyGo {
		// Try to find wasm_exec.js based on the location of tinygo;
		// e.g. if tinygo is at /usr/bin/tinygo, we look under
		// /usr/lib/tinygo and other similar directories.
		tinygoExe, err := exec.LookPath("tinygo")
		if err != nil {
			return "", fmt.Errorf("can't find tinygo executable: %w", err)
		}
		prefix := filepath.Dir(filepath.Dir(tinygoExe))
		candidates := []string{"/lib", "/lib32", "/lib64", "/share"}
		suffix := "/tinygo/targets/wasm_exec.js"
		for _, c := range candidates {
			path := prefix + c + suffix
			if _, err := os.Stat(path); err == nil {
				return path, nil
			}
		}
		return "", fmt.Errorf("failed to find wasm_exec.js")
	}
	// Regular go toolchain
	cmd := exec.Command("go", "env", "GOROOT")
	cmd.Env = append(cmd.Env, os.Environ()...)
	goroot, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not determine GOROOT: %v", err)
	}
	path := strings.TrimSpace(string(goroot)) + "/misc/wasm/wasm_exec.js"
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("could not stat %q: %v", path, err)
	}
	return path, nil
}

func buildWebui(r *BuildRecord, cfg Config) error {
	const (
		tmpPath   = "_build/webui.wasm"
		finalPath = "internal/server/embed/webui.wasm"
		srcDir    = "./cmd/webui"
	)

	wasmExecSrc, err := findWasmExecJs(cfg)
	if err != nil {
		return err
	}

	// Build the webassembly binary:
	log.Println("Building wasm binary")
	if cfg.TinyGo {
		err := runInDir(".",
			"tinygo", "build",
			"-target", "wasm",
			"-panic", "trap",
			"-no-debug",
			"-o="+tmpPath,
			srcDir)
		if err != nil {
			return err
		}
	} else {
		// Use the tsandard go toolchain.
		cmd := exec.Command("go", "build", "-o", tmpPath, srcDir)
		cmd.Env = append(cmd.Env, os.Environ()...)
		cmd.Env = append(cmd.Env, "GOOS=js", "GOARCH=wasm")
		err := withMyOuts(cmd).Run()
		if err != nil {
			return err
		}
	}
	if !r.IsModified(tmpPath) {
		return nil
	}
	chkfatal(r.RecordFile(tmpPath))
	runInDir(".", "du", "-hs", tmpPath)
	chkfatal(copyFile(finalPath, tmpPath))
	return copyFile("internal/server/embed/wasm_exec.js", wasmExecSrc)
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

func buildGo(r *BuildRecord) error {
	buildCapnp(r)

	err := buildWebui(r, readConfig())
	if err != nil {
		return err
	}
	exes := []struct {
		name   string
		static bool
	}{
		{"sandstorm-import-tool", false},
		{"tempest", false},
		{"tempest-make-user", false},
		{"tempest-grain-agent", true},
		{"test-app", true},
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
	cmd := exec.Command("go", "build", "-v", "-o", "_build/"+name, "./cmd/"+name)
	cmd.Env = append(cmd.Env, os.Environ()...)
	if static {
		cmd.Env = append(cmd.Env, "CGO_ENABLED=0")
	} else {
		cmd.Env = append(cmd.Env, "CGO_ENABLED=1")
	}
	return withMyOuts(cmd).Run()
}

func buildTestSpk() error {
	return runInDir("cmd/test-app",
		"spk", "pack",
		"--keyring", "./sandstorm-keyring",
		"../../_build/test-app.spk",
	)

}

// Run configure if its outputs aren't already present.
func maybeConfigure() {
	_, errJson := os.Stat("./config.json")
	if errJson == nil {
		// Config is already present; we're done.
		return
	}
	log.Println("'configure' has not been run; running with default options.")
	run("configure")
}

func run(args ...string) {
	switch args[0] {
	case "build":
		maybeConfigure()
		chkfatal(os.MkdirAll("_build", 0755))
		r := GetBuildRecord()
		buildConfig(r)
		chkfatal(buildC())
		chkfatal(buildGo(r))
		r.Save()
	case "test-app":
		run("build")
		chkfatal(buildTestSpk())
	case "export-import":
		maybeConfigure()
		run("build")
		exportImport(readConfig())
	case "configure":
		cfg := &Config{}
		cfg.ParseFlags(args, "configure", flag.ExitOnError)
		jsonData, err := json.MarshalIndent(cfg, "", "  ")
		chkfatal(err)
		chkfatal(os.WriteFile(
			"./config.json",
			jsonData,
			0600))
	case "install":
		run("build")
		c := readConfig()
		installExe(c, "tempest", c.Bindir, "cap_net_bind_service+ep")
		installExe(c, "tempest-sandbox-launcher", c.Libexecdir+"/tempest",
			"cap_sys_admin,cap_net_admin,cap_mknod+ep")
		installExe(c, "tempest-grain-agent", c.Libexecdir+"/tempest", "")
		chkfatal(os.MkdirAll(c.Localstatedir+"/sandstorm/mnt", 0755))
	case "dev":
		run("install")
		c := readConfig()
		cmd := withMyOuts(
			exec.Command("sudo",
				"--preserve-env",
				"-u", c.User,
				"-g", c.Group,
				c.Bindir+"/tempest",
			),
		)
		cmd.Env = os.Environ()
		chkfatal(cmd.Run())
	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", args[0])
		os.Exit(1)
	}
}

func readConfig() Config {
	var c Config
	data, err := os.ReadFile("config.json")
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

	if stamp == sig.Stamp {
		return false
	}
	hash, err := HashFile(path)
	if err != nil {
		return true
	}
	return !bytes.Equal(hash, sig.Hash)
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

func exportImport(cfg Config) {
	dbPath := cfg.Localstatedir + "/sandstorm/sandstorm.sqlite3"
	chkfatal(runInDir(".", "_build/sandstorm-import-tool", "export"))
	chkfatal(os.Remove(dbPath))
	chkfatal(runInDir(".", "_build/sandstorm-import-tool", "import"))
	chkfatal(runInDir(".", "chown", cfg.User+":"+cfg.Group, dbPath))
}
