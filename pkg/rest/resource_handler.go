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
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(res)
	if !res.Check() || err != nil {
		http.Error(w, "Resource contains invalid data", http.StatusBadRequest)
		return
	}

	if resource.GlobalPool.Contains(res.Name) {
		http.Error(w, "Such resource is already in the cluster", http.StatusBadRequest)
		return
	}

	res.AddKey()
	resource.GlobalPool.Add(*res)
	w.WriteHeader(http.StatusOK)
	logger.Infof("Resource %s was successfully added", res.Name)
}

func ShowResources(w http.ResponseWriter, _ *http.Request) {
	logger.Info("handling resources list request")

	s := struct {
		Res []string `json:"resources"`
	}{}

	l := resource.GlobalPool.List()
	for _, i := range l {
		if i != nil {
			s.Res = append(s.Res, i.(string))
		}
	}

	data, err := json.Marshal(s)
	if err != nil {
		http.Error(w, "Can't marshall resource info "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	logger.Info("resource list have been sent")
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
	w.WriteHeader(http.StatusOK)
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
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(res)
	if !res.Check() || err != nil {
		http.Error(w, "Resource contains invalid data", http.StatusBadRequest)
		return
	}
	res.AddKey()
	resource.GlobalPool.Change(*res)
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
	logger.Infof("Resource %s was successfully removed", resName)
}
