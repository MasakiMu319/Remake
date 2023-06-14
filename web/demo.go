package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is foo")
}

//func main() {
//	http.HandleFunc("/bar", func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Fprint(writer, "Hello, %q", html.EscapeString(request.URL.Path))
//	})
//	log.Fatal(http.ListenAndServe(":8080", foo))
//}
