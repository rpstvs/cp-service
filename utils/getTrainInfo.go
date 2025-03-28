package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

const baseURL = "https://www.cp.pt/sites/spring/station/trains/train?trainId="

func GetTrainInfo(numworkers int) {

	var GeneralTrainInfo []TrainInfo
	data, err := os.ReadFile("trainNumbers.json")

	if err != nil {
		log.Println(err)
		return
	}

	var trainNumbers []int

	err = json.Unmarshal(data, &trainNumbers)

	if err != nil {
		log.Println(err)
		return
	}

	results := make(chan TrainInfo, len(trainNumbers))
	var wg sync.WaitGroup
	jobs := make(chan int, len(trainNumbers))

	for w := 1; w <= 30; w++ {
		go func() {

			for trainId := range jobs {
				RequestWorker(trainId, results, &wg)
			}

		}()
	}

	wg.Add(len(trainNumbers))

	for _, trainID := range trainNumbers {
		jobs <- trainID
	}

	close(jobs)

	wg.Wait()
	close(results)

	for info := range results {
		if info.Status == "" || info.Status == "COMPLETED" {
			continue
		}
		GeneralTrainInfo = append(GeneralTrainInfo, info)
	}

	updatedInfo := LiveStatus{
		Trains:     GeneralTrainInfo,
		Updated_at: time.Now().Unix(),
	}
	log.Printf("Finished fetching Train info, there are %d trains in transit \n", len(GeneralTrainInfo))
	WriteLiveStatusToFile(updatedInfo, "LiveInfo.json")

}

func RequestWorker(trainId int, results chan<- TrainInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(200 * time.Millisecond)
	reqURL := fmt.Sprintf(baseURL + strconv.Itoa(trainId))
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	var info TrainInfo

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&info)
	if err != nil {
		log.Printf("Nothing to decode on trainid %d \n", trainId)
		return
	}

	results <- info

}
