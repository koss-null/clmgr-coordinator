package rest

import (
	"encoding/json"
	"github.com/google/logger"
	"myproj.com/clmgr-coordinator/pkg/cluster"
	"net/http"
)

func ConfigureCluster(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling configure resource request")

	conf := new(cluster.Config)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(conf)
	if !conf.Check() || err != nil {
		http.Error(w, "Cluster info contains invalid data", http.StatusBadRequest)
		return
	}
	cluster.Current.AddConfig(conf)
	w.WriteHeader(http.StatusOK)
	logger.Info("Cluster was successfully configured")
}
