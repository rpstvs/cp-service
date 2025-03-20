package server

import "net/http"

func NewServer(port string) http.Server {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/", RetrieveTrainsInfo)

	return http.Server{
		Handler: serverMux,
		Addr:    ":" + port,
	}
}
