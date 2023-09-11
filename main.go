package main

import (
	"crudone/config"
	"crudone/controllers/categoryController"
	"crudone/controllers/homeController"
	"log"
	"net/http"
)


func main() {

	// panggil ConnectDB 
	config.ConnectDB()

	// HomePage
	http.HandleFunc("/homeIndex", homeController.Welcome)

	// Category 
	http.HandleFunc("/category/Index", categoryController.Index)
	// http.HandleFunc("/category/Add", categoryController.Welcome)
	// http.HandleFunc("/category/Edit", categoryController.Welcome)
	// http.HandleFunc("/category/Delete", categoryController.Welcome)






	// Connect ke Server
	// ketika menjalankan server, jangan pakai fmt.print, tetapi pakai log.println
	log.Println("Server Running On Port 1378")
	http.ListenAndServe(":1378", nil )
}