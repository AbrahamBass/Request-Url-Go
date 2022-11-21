package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo")
	})

	http.ListenAndServe(":3000", nil)
}
