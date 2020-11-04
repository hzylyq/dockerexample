package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handle)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	resp := "You've hit" + hostname + "\n"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
