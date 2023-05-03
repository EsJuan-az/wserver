package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, request *http.Request){
	fmt.Fprintf(w, "Ruta base")
}
func HandleHome(w http.ResponseWriter, request *http.Request){
	fmt.Fprintf(w, "Ruta home")
}