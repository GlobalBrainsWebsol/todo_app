package controllers

import (
	"config"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			users(w, r)
		case "POST":
			userSave(w, r)
		default:
			http.Error(w, "Unsupported request method.", http.StatusMethodNotAllowed)
		}
	})

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
