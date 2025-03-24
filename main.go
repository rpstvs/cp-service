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
			utils.GetTrainInfo(30)
			time.Sleep(30 * time.Second)

		}

	}()

	server := server.NewServer("8080")

	log.Println("Serving on port: ", server.Addr)
	log.Fatal(server.ListenAndServe())

}
