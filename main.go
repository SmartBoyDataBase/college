package main

import (
	"log"
	"net/http"
	"sbdb-college/handler"
)

func main() {
	http.HandleFunc("/ping", handler.PingPongHandler)
	http.HandleFunc("/college", handler.CollegeHandler)
	http.HandleFunc("/colleges", handler.CollegesHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
