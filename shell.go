package main

import "net/http"

func shellHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}