package browsermain

import "syscall/js"

func navigate(newHash string) {
	js.Global().Get("location").Set("hash", newHash)
}
