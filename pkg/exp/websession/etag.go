package websession

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type etag struct {
	Value string
	Weak  bool
}

func parseETagList(s string) (tags []etag, err error) {
	var tag etag
	for {
		tag, s, err = parseETag(s)
		if err == io.EOF {
			return tags, nil
		}
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
		s = strings.TrimSpace(s)
		if s == "" {
			return tags, nil
		}
		if s[0] != ',' {
			return nil, fmt.Errorf("unexpected character in etag list: %c", s[0])
		}
		s = s[1:]
	}
}

func parseETag(s string) (result etag, rest string, err error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return etag{}, "", io.EOF
	}
	if strings.HasPrefix(s, "W/") {
		result.Weak = true
		s = s[2:]
	}
	s = strings.TrimSpace(s)
	if s == "" || s[0] != '"' {
		return etag{}, s, errors.New("unquoted etag value")
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '"' {
			result.Value = strings.TrimSpace(s[:i+1])
			return result, s[i+1:], nil
		}
		if s[i] == '\\' {
			i++
		}
	}
	return etag{}, s, errors.New("no terminating quote in etag value")
}
