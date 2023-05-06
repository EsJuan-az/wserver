package main

import (
	"net/http"
)

type Router struct{
	rules map[string]map[string]http.HandlerFunc
}
func NewRouter()*Router{
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}
func (r *Router)FindHandler(path, method string) (http.HandlerFunc,bool,  bool){
	methods, existPath := r.rules[path]
	handler, methodAllowed := methods[method]
	return handler, existPath, methodAllowed
}
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request){
	handler, existPath, methodAllowed := r.FindHandler(request.URL.Path, request.Method)
	if !existPath{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodAllowed{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}
