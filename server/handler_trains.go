package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rpstvs/cp-realtime/utils"
)

func RetrieveTrainsInfo(w http.ResponseWriter, r *http.Request) {
	var LiveInfo utils.LiveStatus

	jsonFile, err := os.Open("LiveInfo.json")
	if err != nil {
		fmt.Println(err)
	}

	decoder := json.NewDecoder(jsonFile)

	err = decoder.Decode(&LiveInfo)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("deu joia")
	respondWithJson(w, http.StatusOK, LiveInfo)
}
