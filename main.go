package main

import "net/http"

func main() {
	http.HandleFunc("/user", makeHttpHandler(handleGetUserByID))
	http.ListenAndServe(":3000", nil)
}
