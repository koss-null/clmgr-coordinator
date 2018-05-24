package rest

import (
	"encoding/json"
	"github.com/google/logger"
	"myproj.com/clmgr-coordinator/pkg/cluster"
	"net/http"
)

func ConfigureCluster(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling configure resource request")

	cl := cluster.New()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(cl)
	if !cl.Check() || err != nil {
		http.Error(w, "Cluster info contains invalid data", http.StatusBadRequest)
		return
	}
	cluster.Current = cl
	w.WriteHeader(http.StatusOK)
	logger.Info("Cluster was successfully configured")
}
