package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	db, err := NewPostgreSQL()
	if err != nil {
		log.Fatalf("Could not initialize database connection %s", err)
	}
	defer db.Close()

	redis, err := NewRedis()

	if err != nil {
		log.Fatalf("Could not initialize Redis client %s", err)
	}

	router := mux.NewRouter()

	renderJSON := func(w http.ResponseWriter, val interface{}, statusCode int) {
		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(val)
	}

	router.HandleFunc("/names/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		val, err := redis.GetName(r.Context(), id)
		if err == nil {
			renderJSON(w, &val, http.StatusOK)
			return
		}

		name, err := db.FindByNConst(id)
		if err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		_ = redis.SetName(r.Context(), name)

		renderJSON(w, &name, http.StatusOK)
	})

	fmt.Println("Starting Server at port :8080")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
