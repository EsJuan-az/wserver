package wserver

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

// func CheckAuth() Middleware {
// 	return func(f http.HandlerFunc) http.HandlerFunc {
// 		return func(w http.ResponseWriter, r *http.Request) {
// 			flag := true
// 			fmt.Println("Checking Authentiction (1)")
// 			if flag {
// 				f(w, r)
// 			} else {
// 				return
// 			}
// 		}
// 	}
// }
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func(){
				methodColors := map[string]func(format string, a ...interface{})string{
					"GET": color.HiGreenString,
					"POST": color.HiMagentaString,
					"PUT": color.HiBlueString,
					"DELETE": color.HiRedString,
				}
				log.Printf("%s%s",methodColors[r.Method](r.Method),r.URL.Path)
			}()
			f(w, r)
		}
	}
}
