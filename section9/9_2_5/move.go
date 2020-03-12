package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Ikuhiroo"))
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
