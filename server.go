package main

import "net/http"

func main() {
	http.HandleFunc("/time", getTime)
	err := http.ListenAndServe(":8795", nil)
	handleError(err)
}

func getTime(w http.ResponseWriter, r *http.Request) {

}

func handleError(err error) {

}
