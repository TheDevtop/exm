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
	} else if re, err := rec.Generate(reqForm.Regex); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if results := eng.Search(re, streamer); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Results: results})
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
	} else if re, err := rec.Generate(reqForm.Regex); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if results := eng.Replace(re, streamer, reqForm.Mapping); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Results: results})
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
	} else if re, err := rec.Generate(reqForm.Regex); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if results := eng.MapReduce(re, streamer, reqForm.Mapping); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Results: results})
		return
	}
}
