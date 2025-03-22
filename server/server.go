package server

import "net/http"

func NewServer(port string) http.Server {
	serverMux := http.NewServeMux()
	cors := middlewareCors(serverMux)
	serverMux.HandleFunc("GET /api/liveinfo", RetrieveTrainsInfo)

	return http.Server{
		Handler: cors,
		Addr:    ":" + port,
	}
}
