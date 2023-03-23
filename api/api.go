package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TheDevtop/exm/api/forms"
	"github.com/TheDevtop/exm/eng"
	"github.com/TheDevtop/exm/rec"
	"github.com/TheDevtop/exm/sti"
	"github.com/TheDevtop/go-probes"
	"github.com/TheDevtop/tpjson"
)

func apiPing(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "pong") }

func apiSearch(w http.ResponseWriter, r *http.Request) {
	pb := probes.NewLogProbe("api.apiSearch", os.Stderr)
	reqForm := new(forms.SearchForm)

	if err := tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if streamer, err := sti.Stream(reqForm.Object); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if re, err := rec.Receive(reqForm.Regex); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Results: eng.Search(re, streamer)})
		return
	}
}

func apiReplace(w http.ResponseWriter, r *http.Request) {
	pb := probes.NewLogProbe("api.apiReplace", os.Stderr)
	reqForm := new(forms.ReplaceForm)

	if err := tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if streamer, err := sti.Stream(reqForm.Object); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if re, err := rec.Receive(reqForm.Regex); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Results: eng.Replace(re, streamer, reqForm.Mapping)})
		return
	}
}

func apiMapReduce(w http.ResponseWriter, r *http.Request) {
	pb := probes.NewLogProbe("api.apiMapReduce", os.Stderr)
	reqForm := new(forms.ReplaceForm)

	if err := tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if streamer, err := sti.Stream(reqForm.Object); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if re, err := rec.Receive(reqForm.Regex); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Results: eng.MapReduce(re, streamer, reqForm.Mapping)})
		pb.Probe(err.Error())
		return
	}
}

func apiReduce(w http.ResponseWriter, r *http.Request) {
	pb := probes.NewLogProbe("api.apiReduce", os.Stderr)
	reqForm := new(forms.ObjectForm)

	if err := tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if streamer, err := sti.Stream(reqForm.Object); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Results: eng.Reduce(streamer)})
		return
	}
}
