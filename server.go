package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type Server struct{
	port string
	router *Router
}
func NewServer(port string)*Server{
	return &Server{
		port:port,
		router: NewRouter(),
	}
}
func (s *Server)Handle(path string, handler http.HandlerFunc){
	s.router.rules[path] = handler
}
func AddMiddlewares(f http.HandlerFunc, middlewares ...Middleware)http.HandlerFunc{{
	for _,m := range middlewares{
		f = m(f)
	}
	return f
}}
func (s *Server)Listen()error{
	http.Handle("/", s.router)
	fmt.Printf("Running on Port %s with status: ", s.port)
	color.New(color.FgGreen).Printf("*\n")
	err := http.ListenAndServe(s.port, nil)
	return err
}