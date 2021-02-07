package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("calc.html")
	fmt.Fprint(w, string(body))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
