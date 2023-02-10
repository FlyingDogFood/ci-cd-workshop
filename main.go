package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/flyingdogfood/ci-cd-workshop/pkg/hello"
)

func main() {
	http.HandleFunc("/", getHello)
	http.ListenAndServe(":8080", nil)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	language := r.URL.Query().Get("language")

	message := hello.Hello(name, language)

	fmt.Printf("got request, Name: %s, Language: %s\n", name, language)
	io.WriteString(w, message)
}
