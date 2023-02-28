package main

import "net/http"

func main() {
	http.HandleFunc("/time", GetTime)
	err := http.ListenAndServe(":8795", nil)
	handleError(err)
}
