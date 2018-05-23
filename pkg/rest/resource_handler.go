package rest

import (
	"encoding/json"
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"myproj.com/clmgr-coordinator/pkg/resource"
	"net/http"
)

func AddResource(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling adding resource request")

	res := new(resource.Resource)
	var data []byte
	r.Body.Read(data)
	json.Unmarshal(data, res)
	if !res.Check() {
		http.Error(w, "Resource contains invalid data", http.StatusBadRequest)
		return
	}

	if !resource.GlobalPool.Contains(res.Name) {
		http.Error(w, "No such resource in the cluster", http.StatusBadRequest)
		return
	}

	res.AddKey()
	resource.GlobalPool.Add(*res)
	logger.Infof("Resource %s was successfully added", res.Name)
}

func ShowResource(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling resource info request")

	params := mux.Vars(r)
	resName, ok := params["name"]
	if !ok || !resource.GlobalPool.Contains(resName) {
		http.Error(w, "No such resource in the cluster", http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(resource.GlobalPool.Get(resName))
	if err != nil {
		http.Error(w, "Can't marshall resource info "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
	logger.Info("resource info have been sent")
}

func ConfigureResource(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling configure resource request")

	params := mux.Vars(r)
	resName, ok := params["name"]
	if !ok || !resource.GlobalPool.Contains(resName) {
		http.Error(w, "No such resource in the cluster", http.StatusBadRequest)
		return
	}

	res := new(resource.Resource)
	var data []byte
	r.Body.Read(data)
	json.Unmarshal(data, res)
	if !res.Check() {
		http.Error(w, "Resource contains invalid data", http.StatusBadRequest)
		return
	}
	res.AddKey()
	resource.GlobalPool.Change(*res)
	logger.Infof("Resource %s was successfully configured", resName)
}

func RemoveResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resName, ok := params["name"]
	if !ok || !resource.GlobalPool.Contains(resName) {
		http.Error(w, "No such resource in the cluster", http.StatusBadRequest)
		return
	}

	resource.GlobalPool.Remove(resName)
	logger.Infof("Resource %s was successfully removed", resName)
}
