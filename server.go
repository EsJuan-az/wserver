package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}
func (s *Server) Handle(method, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist{
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	fmt.Printf("Running on Port %s with status: ", s.port)
	color.New(color.FgHiGreen).Printf("*\n")
	err := http.ListenAndServe(s.port, nil)
	return err
}
