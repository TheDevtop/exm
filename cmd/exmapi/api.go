package main

import (
	"fmt"
	"net/http"

	"github.com/TheDevtop/exm/cmd/exmapi/engine"
	"github.com/TheDevtop/exm/lib"
	"github.com/TheDevtop/tpjson"
)

const (
	urlSearch_object = "/search/object" // Search into object with regex, return list with matches
	urlSearch_global = "/search/global" // Search into all objects with regex, return list with objects
	urlIndex_object  = "/index/object"  // Return the dictionary of object
	urlIndex_global  = "/index/global"  // Return list of known source objects (config)
	urlMeta_object   = "/meta/object"   // Return metadata of object
	urlPing          = "/ping"          // Ping pong
)

func apiSearch_object(w http.ResponseWriter, r *http.Request) {
	var (
		reqForm = new(lib.FormRequest)
		results []string
		err     error
	)

	if err = tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, lib.FormResult{Error: err.Error()})
		return
	}
	if results, err = engine.SearchObject(reqForm.Regex, reqForm.Object); err != nil {
		tpjson.SendJSON(w, lib.FormResult{Error: err.Error()})
		return
	}
	tpjson.SendJSON(w, lib.FormResult{Count: len(results), Results: results})
}

func apiSearch_global(w http.ResponseWriter, r *http.Request) {
	var (
		reqForm = new(lib.FormRequest)
		results []string
		err     error
	)

	if err = tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, lib.FormResult{Error: err.Error()})
		return
	}
	if results, err = engine.SearchGlobal(reqForm.Regex); err != nil {
		tpjson.SendJSON(w, lib.FormResult{Error: err.Error()})
		return
	}
	tpjson.SendJSON(w, lib.FormResult{Count: len(results), Results: results})
}

func apiIndex_object(w http.ResponseWriter, r *http.Request) {
	var (
		reqForm = new(lib.FormRequest)
		results []string
		err     error
	)

	if err = tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, lib.FormResult{Error: err.Error()})
		return
	}
	if results, err = engine.IndexObject(reqForm.Object); err != nil {
		tpjson.SendJSON(w, lib.FormResult{Error: err.Error()})
		return
	}
	tpjson.SendJSON(w, lib.FormResult{Count: len(results), Results: results})
}

func apiIndex_global(w http.ResponseWriter, r *http.Request) {
	var results = engine.IndexGlobal()
	tpjson.SendJSON(w, lib.FormResult{Count: len(results), Results: results})
}

func apiMeta_object(w http.ResponseWriter, r *http.Request) {
	var (
		reqForm = new(lib.FormRequest)
		results lib.FormMetadata
		err     error
	)

	if err = tpjson.ReceiveJSON(r, reqForm); err != nil {
		tpjson.SendJSON(w, lib.FormResult{Error: err.Error()})
		return
	}
	if results = engine.MetaObject(reqForm.Object); results.Error != "" {
		tpjson.SendJSON(w, lib.FormResult{Error: results.Error})
		return
	}
	tpjson.SendJSON(w, results)
}

func apiPing(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "pong") }
