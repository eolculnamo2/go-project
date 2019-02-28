package controllers

import (
		"log"
		"net/http"
		"html/template"
		"encoding/json"
		)

//data structures
type user struct {
	User string
	password string
}
	
func index_handler(w http.ResponseWriter, r *http.Request) {
	t,_ := template.ParseFiles("public/index.html")
	t.Execute(w, nil)
}

func login_handler(w http.ResponseWriter, req*http.Request) {
	if req.Method == "POST" {
		decoder := json.NewDecoder(req.Body)
		var user user
		err := decoder.Decode(&user)
		if err != nil {
			panic(err)
		}
		log.Println(user.User)
	}
}