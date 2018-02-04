package main

import (
	"github.com/gorilla/sessions"
)

type RpcServer struct {
	Store sessions.Store
}

func (s *RpcServer) ListApps(arg struct{}, ret *[]string) error {
	*ret = make([]string, 0, len(apps))
	for k, _ := range apps {
		*ret = append(*ret, k)
	}
	return nil
}
