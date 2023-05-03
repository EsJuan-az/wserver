package main

import "net/http"

func main(){
	routes := map[string]http.HandlerFunc{
		"/": HandleRoot,
		"/home": HandleHome,
	}
	server := NewServer(":2002")
	for path, handler := range routes{
		server.Handle(path, handler)
	}
	server.Listen()
}