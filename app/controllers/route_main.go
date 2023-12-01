package controllers

import (
	"app/models"
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil) //第2引数はdataを渡す。view側{{.}}で表示する
}

func users(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/templates/users.html")
	if err != nil {
		log.Fatalln(err)
	}
	ul, _ := models.UserLists()
	t.Execute(w, ul)
}

func userSave(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	pass := r.PostFormValue("password")
	u := &models.User{}
	u.Name = name
	u.Email = email
	u.Password = pass
	u.CreateUser()
	http.Redirect(w, r, "/users", http.StatusFound)
}
