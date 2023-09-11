package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

	var DB *sql.DB

func ConnectDB() {

	// nama DB = go-crudOne 
	// Jika ada error di hentikan 
	db, err := sql.Open("mysql", "root:@/go-crudOne")
	if err != nil {
		panic(err)
	}
	
	DB = db
		log.Println("Database Connected")


}
