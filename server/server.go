package server

import "net/http"

func NewServer(port string) http.Server {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("GET /api/liveinfo", RetrieveTrainsInfo)

	return http.Server{
		Handler: serverMux,
		Addr:    ":" + port,
	}
}
