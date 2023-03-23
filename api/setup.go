package api

import "net/http"

const ListenAddr = ":1800"

func Setup() {
	http.HandleFunc("/ping", apiPing)
	http.HandleFunc("/search", apiSearch)
	http.HandleFunc("/replace", apiReplace)
	http.HandleFunc("/map/reduce", apiMapReduce)
}