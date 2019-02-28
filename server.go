package main

import ("net/http"
		"./controllers"
		"./api")

func main() {
	fs := http.FileServer(http.Dir("public"))
  	http.Handle("/public/", http.StripPrefix("/public/", fs))
	controllers.Controller()
	api.ConnectToDb();
	http.ListenAndServe(":8000", nil)
}