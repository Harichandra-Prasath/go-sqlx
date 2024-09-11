package main

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Addr string
}

type Server struct {
	*Config
	DB  *sqlx.DB
	Mux *http.ServeMux
}

func GetServer(c *Config) *Server {
	return &Server{
		c,
		DB_INIT(),
		http.NewServeMux(),
	}
}

func (S *Server) Start() {

	fmt.Println("Server Started Listening on Port", S.Addr)

	err := http.ListenAndServe(S.Addr, S.Mux)
	if err != nil {
		panic(err)
	}
}
