package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// w.Write([]byte("hello jun blog"))

	// jsonS, _ := json.Marshal("hello jun blog")
	// w.Write(jsonS)

	var indD IndexData
	indD.Title = "Jun Blog"
	indD.Desc = "The first Blog"
	indJson, _ := json.Marshal(indD)
	w.Write(indJson)

}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
