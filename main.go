package main

import (
	"os"
	"log"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

func main {
	port := os.Getenv("PORT")
	dsn := os.Getenv("DB_DSN")
	
	if port == "" {log.Fatal("$PORT must be set")}
	if dsn == "" {log.Fatal("$DB_DSN must be set")}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	
	db.AutoMigrate(&GoFooBar{})
	
	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/read/", readHandler)
	http.HandleFunc("/update/", updateHandler)
	http.HandleFunc("/delete/", deleteHandler)

	http.ListenAndServe(":"+port, nil)
}
