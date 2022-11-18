package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type application struct {
	db    *PostgreSQL
	redis *Client
}

func main() {
	app := &application{}

	app.db = InitializePostgres()
	app.redis = InitializeRedis()

	defer app.db.Close()

	fmt.Println("Starting Server at port :8080")

	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
