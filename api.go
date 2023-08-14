package main

import (
	"encoding/json"
	"net/http"
)

// function type for handling API requests.
type apiFunc func(http.ResponseWriter, *http.Request) error

// takes an apiFunc and returns an http.HandlerFunc that handles
// incoming HTTP requests by calling the provided apiFunc
func makeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJson(w, e.Status, e)
				return
			}
			writeJson(w, http.StatusInternalServerError, apiError{Err: "internal server"})
		}
	}
}

// an apiFunc implementation that handles HTTP GET requests for
// retrieving user information by ID.
func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}
	}

	user := User{}

	if !user.Valid {
		return apiError{Err: "User not found", Status: http.StatusForbidden}
	}

	return writeJson(w, http.StatusOK, User{})

}

// writes a JSON-encoded response with the provided status code
// and data to the given http.ResponseWriter.
func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
