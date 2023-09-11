package main

import (
	"crudone/config"
	"crudone/controllers/homeController"
	"log"
	"net/http"
)


func main() {

	// panggil ConnectDB 
	config.ConnectDB()

	// HomePage
	http.HandleFunc("/homeIndex", homeController.Welcome)



	// Connect ke Server
	// ketika menjalankan server, jangan pakai fmt.print, tetapi pakai log.println
	log.Println("Server Running On Port 1378")
	http.ListenAndServe(":1378", nil )
}