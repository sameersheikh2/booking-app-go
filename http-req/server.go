package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("konichiwa")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hello world")) })

	http.ListenAndServe(":5050", nil)
}
