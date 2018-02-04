package main

import (
	"fmt"
	"os"
	"os/exec"
)

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

func buildElm() error {
	cmd := exec.Command("elm-make", "Main.elm", "--yes", "--output", "../static/index.html")
	cmd.Dir = "elm"
	return withMyOuts(cmd).Run()
}

func buildGo() error {
	return buildGoBin("server")
}
func buildGoBin(bin string) error {
	return withMyOuts(exec.Command(
		"go", "build", "-v", "-i", "-o", "bin/"+bin, "./go/cmd/"+bin,
	)).Run()
}

func run(args ...string) {
	switch args[0] {
	case "build":
		elmCh := spawnErr(buildElm)
		goCh := spawnErr(buildGo)
		elmErr := <-elmCh
		goErr := <-goCh
		if elmErr != nil || goErr != nil {
			fmt.Fprintln(os.Stderr, "Errors occured.")
			os.Exit(1)
		}
	case "run":
		run("build")
		fmt.Fprintln(os.Stderr, "Starting server...")
		err := withMyOuts(exec.Command("./bin/server")).Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "run: error:", err)
		}
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
