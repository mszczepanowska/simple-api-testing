package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("access denied"))
}


func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	fmt.Println("Start server")
	http.HandleFunc("/", root)
	http.HandleFunc("/api/status", status)
	http.ListenAndServe(":8080", nil)
}
