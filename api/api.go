package api

import (
	"fmt"
	"net/http"

	"github.com/TheDevtop/exm/api/forms"
	"github.com/TheDevtop/exm/conio"
	"github.com/TheDevtop/exm/eng"

	"github.com/TheDevtop/tpjson"
)

func apiPing(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "pong") }

func apiSearch(w http.ResponseWriter, r *http.Request) {
	const fprobe = "api.apiSearch"
	reqForm := new(forms.SearchForm)

	if err := tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		conio.Probeln(fprobe, err.Error())
		return
	} else if streamer, err := conio.Stream(reqForm.Object); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		conio.Probeln(fprobe, err.Error())
		return
	} else if results, err := eng.Search(reqForm.Regex, streamer); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		conio.Probeln(fprobe, err.Error())
		return
	} else {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Results: results})
		return
	}
}

func Setup() {
	http.HandleFunc("/ping", apiPing)
	http.HandleFunc("/search", apiSearch)
}
