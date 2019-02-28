package controllers

import "net/http"

func Controller() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/login", login_handler)
}