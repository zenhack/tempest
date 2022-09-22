package main

import (
	"mime"
	"strconv"
)

var (
	responseHeaderWhiteList = map[string]headerSanitizer{
		"Content-Type":   compose(onlyOne, mustSatisfy(isMediaType)),
		"Content-Length": compose(onlyOne, mustSatisfy(isNumber)),
	}
)

type headerSanitizer func(hdr []string) []string

func verbatium(want string) func(string) (string, bool) {
	return func(got string) (string, bool) {
		return got, want == got
	}
}

func any(fns ...headerSanitizer) headerSanitizer {
	return func(hdr []string) []string {
		for _, fn := range fns {
			ret := fn(hdr)
			if ret != nil {
				return ret
			}
		}
		return nil
	}
}

func compose(fns ...headerSanitizer) headerSanitizer {
	return func(hdr []string) []string {
		for _, fn := range fns {
			if hdr == nil {
				break
			}
			hdr = fn(hdr)
		}
		return hdr
	}
}

func isMediaType(s string) bool {
	_, _, err := mime.ParseMediaType(s)
	return err == nil
}

func isNumber(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func each(fn func(string) (string, bool)) headerSanitizer {
	return func(hdr []string) []string {
		for i := range hdr {
			var ok bool
			hdr[i], ok = fn(hdr[i])
			if !ok {
				return nil
			}
		}
		return hdr
	}
}

func onlyOne(hdr []string) []string {
	if len(hdr) > 1 {
		return nil
	}
	return hdr
}

func mustSatisfy(pred func(string) bool) headerSanitizer {
	return each(func(hdr string) (string, bool) {
		return hdr, pred(hdr)
	})
}
