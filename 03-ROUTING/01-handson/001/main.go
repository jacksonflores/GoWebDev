package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar/", bar)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "ayo foo just foo'd up this place")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "im all barred up rn homie")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "its jackson")
}
