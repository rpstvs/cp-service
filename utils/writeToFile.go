package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func writeTrainNumberToFile(trainNumber map[int]struct{}, filename string) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	keys := make([]int, 0, len(trainNumber))
	for k := range trainNumber {
		keys = append(keys, k)
	}

	data, err := json.MarshalIndent(keys, " ", " ")
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	_, err = file.Write([]byte(data))

	if err != nil {
		fmt.Println(err)
	}
}

func WriteLiveStatusToFile(LiveStatus LiveStatus, filename string) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(LiveStatus, " ", " ")
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	_, err = file.Write([]byte(data))

	if err != nil {
		fmt.Println(err)
	}
}
