package main

import (
	"log"
	"time"

	"github.com/rpstvs/cp-realtime/server"
	"github.com/rpstvs/cp-realtime/utils"
)

func main() {
	//utils.GetTrainNumbers()

	go func() {

		for {
			time.Sleep(30 * time.Second)
			utils.GetTrainInfo(30)
		}

	}()

	server := server.NewServer("8080")

	log.Println("Serving on port: ", server.Addr)
	log.Fatal(server.ListenAndServe())

}
