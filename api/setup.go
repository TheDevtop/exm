package api

import "net/http"

const ListenAddr = ":1800"

func Setup() {
	http.HandleFunc("/ping", apiPing)

	http.HandleFunc("/search", apiSearch)
	http.HandleFunc("/search/all", apiSearchAll)

	http.HandleFunc("/replace", apiReplace)
	http.HandleFunc("/replace/all", apiReplaceAll)

	http.HandleFunc("/reduce", apiReduce)
	http.HandleFunc("/reduce/all", apiReduceAll)

	http.HandleFunc("/map/reduce", apiMapReduce)
	http.HandleFunc("/map/reduce/all", apiMapReduceAll)
}
