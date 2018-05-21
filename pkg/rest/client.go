package rest

import (
	"encoding/json"
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"myproj.com/clmgr-coordinator/config"
	. "myproj.com/clmgr-coordinator/pkg/common"
	"net/http"
	"strings"
)

type (
	client struct{}

	Client interface {
		Start() error
	}
)

func NewClient() Client {
	return &client{}
}

func HostnameHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling hostname request")
	hn := CreateHostname()
	var s struct {
		Hostname string `json:"hostname"`
	}
	s.Hostname = hn
	json.NewEncoder(w).Encode(&s)
	logger.Info("Returned hostname")
}

func (*client) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/hostname", HostnameHandler).Methods("GET")
	defaultPort := strings.Split(config.Config.CoordinatorAddress, ":")[1]
	logger.Infof("Start listener on port %s", defaultPort)
	http.ListenAndServe(":"+defaultPort, router)
	return nil
}
