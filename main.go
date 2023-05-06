package wserver

// import "net/http"

// func main() {
	// server := NewServer(":2002")



	// routes := map[string]map[string]http.HandlerFunc{
	// 	"/": {
	// 		"GET": HandleRoot(CheckAuth(), Logging()),
	// 		"POST": HandleUserPost(Logging()),
	// 	},
	// 	"/home": {
	// 		"GET": HandleHome(Logging()),
	// 	},
	// }

	// // AddMiddleware(HandleRoot, CheckAuth(), Logging())
	// // HandleHome
	// for path, methods := range routes {
	// 	for method, handler := range methods {
	// 		server.Handle(method, path, handler)
	// 	}
	// }
	// server.Listen()
// }
