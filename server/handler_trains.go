package server

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rpstvs/cp-realtime/utils"
)

func RetrieveTrainsInfo(w http.ResponseWriter, r *http.Request) {
	var LiveInfo utils.LiveStatus

	jsonFile, err := os.Open("LiveInfo.json")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldnt get Live Info", err)
		return
	}

	decoder := json.NewDecoder(jsonFile)

	err = decoder.Decode(&LiveInfo)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldnt get Live Info", err)
		return
	}

	respondWithJson(w, http.StatusOK, LiveInfo)
}
