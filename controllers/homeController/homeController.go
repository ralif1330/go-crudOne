package homeController

import (
	"net/http"
	"text/template"

)

// diwajibkan dua parameter = Response dan Request
// Respon = W , Request = R
// ini adalah parameter wajib ketika membuat  handler function


func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/homeViews/indexHome.html")
	if err != nil {
		panic(err)
	}

	
	temp.Execute(w, nil)
}