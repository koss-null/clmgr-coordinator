package rest

import (
	"encoding/json"
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"myproj.com/clmgr-coordinator/pkg/node"
	"net/http"
)

func AddLabelHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling hostname request")

	params := mux.Vars(r)
	hostname, ok := params["hostname"]
	if !ok || node.NodePool.Contains(hostname) {
		http.Error(w, "No such resource in the cluster", http.StatusBadRequest)
		return
	}

	s := struct {
		Labels []string `json:"labels,omitempty"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&s)
	if err != nil {
		http.Error(w, "Can't unmarshal labels", http.StatusInternalServerError)
		return
	}

	node.NodePool.AddLabel(hostname, s.Labels)
	logger.Info("Labels was added")
}
