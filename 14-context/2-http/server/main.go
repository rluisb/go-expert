package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		message := "Request processada com sucesso"
		log.Println(message)
		w.Write([]byte(message))
	case <-ctx.Done():
		message := "Request cancelada pelo cliente"
		log.Println(message)
		http.Error(w, message, http.StatusRequestTimeout)
	}

}
