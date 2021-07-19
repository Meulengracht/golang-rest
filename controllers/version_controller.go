package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type serverVersion struct {
	Version string `json:"version"`
}

func GetServerVersion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	version := serverVersion{Version: "1.0.0"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}
