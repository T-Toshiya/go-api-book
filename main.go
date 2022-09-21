package main

import (
	"database/sql"
	"fmt"
	"go-api-book/controllers"
	"go-api-book/routers"
	"go-api-book/services"
	"log"
	"net/http"
	"os"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)
	r := routers.NewRouter(con)

	log.Println("server start at port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
