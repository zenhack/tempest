// `sandstorm-sandbox-agent` runs inside the grain's sandbox, and is the first
// program executed during grain startup. Its file descriptor #3 is a socket
// over which we can speak capnp to the sandstorm server outside the sandbox.
//
// Any APIs available to the grain which don't actually need privileges the grain
// doesn't have should ideally be implemented here; this helps us minimize attack
// surface.
package main

import "fmt"

func main() {
	fmt.Println("Hello, Sandbox!")
}
