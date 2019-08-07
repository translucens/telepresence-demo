package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	backendURL = "http://backend-1:8080"
)

func main() {

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Printf("invalid port specified: %v", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(backendURL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("failed to get a content from backend: %v", err)
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Via", fmt.Sprintf("%s Go-demo-proxy", r.Proto))
		w.WriteHeader(http.StatusOK)
		bytes, err := io.Copy(w, resp.Body)
		if err != nil {
			log.Printf("failed to send a response: %v", err)
			return
		}
		log.Printf("sent response: %d bytes", bytes)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
