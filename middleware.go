package main

import "net/http"

const (
	USERNAME = "fanialfi"
	PASSWORD = "saichiopy"
)

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "something went wrong", http.StatusUnauthorized)
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		http.Error(w, "wrong username/password", http.StatusForbidden)
		return false
	}
	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		http.Error(w, "Only Accept GET method", http.StatusMethodNotAllowed)
		return false
	}
	return true
}
