package categoryController

import (
	categorymodel "crudone/models/categoryModel"
	"net/http"
	"text/template"
)

// panggil function dari folder model
// panggil interface kosong atau any

func Index(w http.ResponseWriter, r *http.Request) {
	categories :=	categorymodel.GetAll()
	data := map[string]any {
		"categories" : categories,
	}

	temp, err := template.ParseFiles("views/categoryView/indexCategory.html")
		if err != nil {
			panic(err)
		}

			temp.Execute(w, data)

}

func Add(w http.ResponseWriter, r *http.Request) {
	
}

func Edit(w http.ResponseWriter, r *http.Request) {
	
}

func Delete(w http.ResponseWriter, r *http.Request) {
	
}