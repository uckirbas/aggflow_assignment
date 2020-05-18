package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/uckirbas/AggFlowAssignment/common"
	"github.com/uckirbas/AggFlowAssignment/models"
)

func SaveBoatEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		var boat models.Boat

		json.Unmarshal(reqBody, &boat)

		if err := boat.Commodity.IsValid(); err != nil {
			log.Print(err)
			return
		}

		log.Print(boat)
		log.Print("**")

		err = env.DB.CreateBoat(&boat)

		if err != nil {
			w.Write([]byte("ERROR: Failed to save Boat!"))
		} else {
			w.Write([]byte("Boat saved successfully!"))
		}

	})
}
