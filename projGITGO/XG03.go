package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Deg usando golang\n")
	})
	log.Fatal(http.ListenAndServe(":8585", nil))
}
