package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
func Embed(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	{
		for i := len(middlewares) - 1; i >= 0; i-- {
			m := middlewares[i]
			f = m(f)
		}
		return f
	}
}

func HandleRoot(middlewares ...Middleware)http.HandlerFunc{
	handler := func(w http.ResponseWriter, request *http.Request){
		fmt.Fprintf(w, "Ruta base")
	}
	return Embed(handler, middlewares...)
}
func HandleHome(middlewares ...Middleware)http.HandlerFunc{
	handler := func(w http.ResponseWriter, request *http.Request){
		fmt.Fprintf(w, "Ruta home")
	}
	return Embed(handler, middlewares...)
}

func HandleUserPost(middlewares ...Middleware)http.HandlerFunc{
	handler := func(w http.ResponseWriter, request *http.Request){
		decoder := json.NewDecoder(request.Body)
		var user User
		err := decoder.Decode(&user)
		if err != nil{
			fmt.Fprintf(w, "error: %v", err)
			return
		}
		response, err := user.toJSON()
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
	return Embed(handler, middlewares...)
}