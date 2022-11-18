package main

import (
	"encoding/json"
	gm "github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
	"redis_postgres_MC/models"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New()

	renderJSON := func(w http.ResponseWriter, val interface{}, statusCode int) {
		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(val)
	}

	redis := app.redis
	db := app.db

	router := gm.NewRouter()

	router.HandleFunc("/names/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := gm.Vars(r)["id"]

		val, err := redis.GetName(r.Context(), id)
		if err == nil {
			renderJSON(w, &val, http.StatusOK)
			return
		}

		name, err := db.FindByNConst(id)
		if err != nil {
			renderJSON(w, &models.Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		_ = redis.SetName(r.Context(), name)

		renderJSON(w, &name, http.StatusOK)
	})

	return standardMiddleware.Append().Then(router)
}
