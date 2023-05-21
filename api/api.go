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

func apiSearchAll(w http.ResponseWriter, r *http.Request) {
	pb := probes.NewLogProbe("api.apiSearchAll", os.Stderr)
	reqForm := new(forms.SearchForm)
	resForm := new(forms.MultiResultForm)
	resForm.Route = r.URL.Path
	resForm.Results = make(map[string][]string)

	if err := tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if re, err := rec.Receive(reqForm.Regex); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if list, err := sti.List(); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		for _, object := range list {
			if streamer, err := sti.Stream(object); err != nil {
				tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
				pb.Probe(err.Error())
			} else {
				resForm.Results[object] = eng.Search(re, streamer)
			}
		}
		tpjson.SendJSON(w, *resForm)
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

func apiReplaceAll(w http.ResponseWriter, r *http.Request) {
	pb := probes.NewLogProbe("api.apiReplaceAll", os.Stderr)
	reqForm := new(forms.ReplaceForm)
	resForm := new(forms.MultiResultForm)
	resForm.Route = r.URL.Path
	resForm.Results = make(map[string][]string)

	if err := tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if re, err := rec.Receive(reqForm.Regex); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if list, err := sti.List(); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		for _, object := range list {
			if streamer, err := sti.Stream(object); err != nil {
				tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
				pb.Probe(err.Error())
			} else {
				resForm.Results[object] = eng.Replace(re, streamer, reqForm.Mapping)
			}
		}
		tpjson.SendJSON(w, *resForm)
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
		tpjson.SendJSON(w, forms.ListForm{Route: r.URL.Path, Count: eng.MapReduce(re, streamer, reqForm.Mapping)})
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

func apiReduceAll(w http.ResponseWriter, r *http.Request) {
	pb := probes.NewLogProbe("api.apiReduceAll", os.Stderr)
	resForm := new(forms.MultiResultForm)
	resForm.Route = r.URL.Path
	resForm.Results = make(map[string][]string)

	if list, err := sti.List(); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		for _, object := range list {
			if streamer, err := sti.Stream(object); err != nil {
				tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
				pb.Probe(err.Error())
			} else {
				resForm.Results[object] = eng.Reduce(streamer)
			}
		}
		tpjson.SendJSON(w, *resForm)
		return
	}
}

func apiListAll(w http.ResponseWriter, r *http.Request) {
	// Allocate objects
	pb := probes.NewLogProbe("api.apiListAll", os.Stderr)
	reqForm := new(forms.SearchForm)
	resForm := new(forms.ListForm)

	// Initialize response form
	resForm.Route = r.URL.Path
	resForm.Count = 0
	resForm.Objects = make([]string, 0)

	if err := tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if re, err := rec.Receive(reqForm.Regex); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else if list, err := sti.List(); err != nil {
		tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
		pb.Probe(err.Error())
		return
	} else {
		for _, object := range list {
			if streamer, err := sti.Stream(object); err != nil {
				tpjson.SendJSON(w, forms.ResultForm{Route: r.URL.Path, Error: err.Error()})
				pb.Probe(err.Error())
			} else {
				if eng.Match(re, streamer) {
					resForm.Objects = append(resForm.Objects, object)
					resForm.Count++
				}
			}
		}
		tpjson.SendJSON(w, *resForm)
		return
	}
}
