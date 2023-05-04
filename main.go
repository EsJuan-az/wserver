package main

import "net/http"

func main() {
	server := NewServer(":2002")
	routes := map[string]http.HandlerFunc{
		"/":     AddMiddleware(HandleRoot, CheckAuth(), Logging()),
		"/home": HandleHome,
	}
	for path, handler := range routes {
		server.Handle(path, handler)
	}
	server.Listen()
}
