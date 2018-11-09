package main

import (
	"io"
	"net/http"
)
import (
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:9999", r)
}
func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello World!")
}
