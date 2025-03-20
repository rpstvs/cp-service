package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rpstvs/cp-realtime/server"
)

func main() {
	//utils.GetTrainNumbers()

	//utils.GetTrainInfo(30)

	go func() {

		for {
			time.Sleep(5 * time.Second)
			foo("estou a fazer calls a cada 30 s")
		}

	}()

	server := server.NewServer("8080")

	log.Println("Serving on port: ", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func foo(message string) {
	fmt.Println(message)
}
