package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uckirbas/AggFlowAssignment/common"
)

func UpdateAssignedCaptainEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		captainId := params["captain_id"]
		boatId := params["boat_id"]

		err := env.DB.UpdateAssignedCaptain(captainId, boatId)

		if err != nil {
			w.Write([]byte("ERROR: Failed to Update Assigned Captain!"))
		} else {
			w.Write([]byte("Captain Assigned successfully!"))
		}

	})
}
