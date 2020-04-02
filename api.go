package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func contentTypeJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func respondJSON(resp interface{}, w http.ResponseWriter, r *http.Request) error {
	contentTypeJSON(w, r)
	j := json.NewEncoder(w)
	err := j.Encode(resp)
	return err
}

type API struct {
	Context *Context
}

func routeApi(router *mux.Router, context *Context) {
	api := API{context}
	router.HandleFunc("/api/state", api.getState).Methods("GET")
	router.HandleFunc("/api/state", api.putState).Methods("PUT")
}

func (api *API) getState(w http.ResponseWriter, r *http.Request) {
	respondJSON(api.Context.State, w, r)
}

func (api *API) putState(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	hr, _ := strconv.Atoi(r.FormValue("heartRate"))
	acc, _ := strconv.Atoi(r.FormValue("accuracy"))

	now := time.Now()

	api.Context.State.CurrentAccuracy = acc
	api.Context.State.CurrentHeartRate = hr
	api.Context.State.DataReceivedAt = now

	if api.Context.WriteToCsv {
		go func() {
			api.Context.CsvFile.WriteString(
				fmt.Sprintf(
					"%s,%d,%d\n",
					now.UTC(),
					hr,
					acc,
				),
			)
		}()
	}

	respondJSON(api.Context.State, w, r)
}
