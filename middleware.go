package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware{
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := false
			fmt.Println("Checking Authentiction")
			if flag{
				f(w, r)
			}else{
				return 
			}
		}
	}
}