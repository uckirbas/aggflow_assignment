package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uckirbas/AggFlowAssignment/common"
)

func GetCaptainProfileEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		boatId := params["boat_id"]
		boat, err := env.DB.GetBoat(boatId)

		if err != nil {
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(boat)

	})
}
