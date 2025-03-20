package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetTrainNumbers() {
	trainNumbers := map[int]struct{}{}
	for _, stationId := range stationList {
		url := fmt.Sprintf("https://www.cp.pt/sites/spring/station/trains?stationId=%s", stationId)
		resp, err := http.Get(url)
		log.Printf("Getting train numbers from %s \n", stationId)
		if err != nil {
			log.Println("No request err:", err)
		}

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var StationInfo TrainStationsInfo
		err = decoder.Decode(&StationInfo)

		if err != nil {
			log.Println("Couldnt decode response err: ", err)
		}

		if len(StationInfo) == 0 {
			continue
		}

		for _, Trains := range StationInfo {
			_, ok := trainNumbers[Trains.TrainNumber]
			if ok {
				continue
			}
			log.Printf("adding train %d", Trains.TrainNumber)
			trainNumbers[Trains.TrainNumber] = struct{}{}
		}

	}
	writeToFile(trainNumbers, "trainNumbers.json")
}
