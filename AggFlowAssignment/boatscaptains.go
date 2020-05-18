package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uckirbas/AggFlowAssignment/common"
	"github.com/uckirbas/AggFlowAssignment/common/datastore"
	"github.com/uckirbas/AggFlowAssignment/endpoints"
)

const (
	WEBSERVERPORT = ":8443"
)

func main() {

	db, err := datastore.NewDatastore(datastore.POSTGRESQL, "postgres://postgres:1234@localhost/aggflow_assignment?sslmode=disable")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	env := common.Env{DB: db}

	r := mux.NewRouter()

	r.Handle("/restapi/get-captain", endpoints.GetCaptainProfileEndpoint(&env)).Methods("GET")
	r.Handle("/restapi/get-boat/{boat_id}", endpoints.GetBoatProfileEndpoint(&env)).Methods("GET")

	r.Handle("/restapi/post-captain", endpoints.SaveCaptainEndpoint(&env)).Methods("POST")

	r.Handle("/restapi/save-boat", endpoints.SaveBoatEndpoint(&env)).Methods("POST")
	r.Handle("/restapi/update-assigned-captain/cpt/{captain_id}/boat/{boat_id}", endpoints.UpdateAssignedCaptainEndpoint(&env)).Methods("PUT")

	http.ListenAndServe(WEBSERVERPORT, r)

}
