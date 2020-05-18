package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/uckirbas/AggFlowAssignment/common"
	"github.com/uckirbas/AggFlowAssignment/models"
)

func SaveCaptainEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		var captain models.Captain
		json.Unmarshal(reqBody, &captain)

		err = env.DB.CreateCaptain(&captain)

		if err != nil {
			w.Write([]byte("ERROR: Failed to save Captain!"))
		} else {
			w.Write([]byte("Captain saved successfully!"))
		}

	})
}
