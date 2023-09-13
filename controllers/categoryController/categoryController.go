package categoryController

import (
	"crudone/entities"
	categorymodel "crudone/models/categoryModel"
	"net/http"
	"strconv"
	"text/template"
	"time"

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

		if r.Method == "GET" {
			temp, err := template.ParseFiles("views/categoryView/createCategory.html")
				if err != nil {
					panic(err) 
				}
				temp.Execute(w,nil)
		}

		if r.Method == "POST" {
			var category entities.Category
			
			category.Name = r.FormValue("name")
			category.CreatedAt = time.Now()
			category.UpdatedAt = time.Now() 

			if ok := categorymodel.Create(category); !ok {
			temp, _ := 	template.ParseFiles("views/categoryView/createCategory.html")
				temp.Execute(w, nil)

			}
				
			http.Redirect(w, r, "/category/Index", http.StatusSeeOther)
		}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/categoryView/editCategory.html")
		if err != nil {
			panic(err) 
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category := categorymodel.Detail(id)
			data := map[string]any{
				"category" : category,	
			}
			
			temp.Execute(w, data)
	}	
}

func Delete(w http.ResponseWriter, r *http.Request) {
	
}