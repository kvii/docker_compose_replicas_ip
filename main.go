package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := os.Getenv("name")
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "name: %s, hostname: %s\n", name, hostname)
	})
	http.ListenAndServe(":9090", nil)
}
