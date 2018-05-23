package rest

import (
	"encoding/json"
	"github.com/google/logger"
	. "myproj.com/clmgr-coordinator/pkg/common"
	"net/http"
)

func HostnameHandler(w http.ResponseWriter, _ *http.Request) {
	logger.Info("handling hostname request")
	hn := CreateHostname()
	var s struct {
		Hostname string `json:"hostname"`
	}
	s.Hostname = hn
	json.NewEncoder(w).Encode(&s)
	logger.Info("Returned hostname")
}
