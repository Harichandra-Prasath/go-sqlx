package main

import (
	"fmt"
	"net/http"
)

type Config struct {
	Addr string
}

type Server struct {
	*Config
}

func GetServer(c *Config) *Server {
	return &Server{
		c,
	}
}

func (S *Server) Start() {

	mux := http.NewServeMux()
	SetRoutes(mux)

	fmt.Println("Server Started Listening on Port", S.Addr)

	err := http.ListenAndServe(S.Addr, mux)
	if err != nil {
		panic(err)
	}
}
