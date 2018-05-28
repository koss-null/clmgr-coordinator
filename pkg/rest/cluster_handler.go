package rest

import (
	"encoding/json"
	"fmt"
	"github.com/google/logger"
	"myproj.com/clmgr-coordinator/pkg/cluster"
	"net/http"
)

func ConfigureCluster(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling configure cluster request")

	conf := new(cluster.Config)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(conf)
	if !conf.Check() || err != nil {
		http.Error(w, fmt.Sprintf("Cluster info contains invalid data, err: %s", err.Error()), http.StatusBadRequest)
		return
	}
	cluster.Current.AddConfig(conf)
	w.WriteHeader(http.StatusOK)
	logger.Info("Cluster was successfully configured")
}

func ShowCluster(w http.ResponseWriter, _ *http.Request) {
	logger.Info("handling show cluster request")

	conf := cluster.Current.GetConfig()
	data, err := json.Marshal(conf)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cluster info can't be marshaled, err: %s", err.Error()), http.StatusBadRequest)
		return
	}

	w.Write(data)
	logger.Info("Cluster was successfully configured")
}
