package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/alexellis/go-modules-test/github-go-tester/pkg"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	org := os.Getenv("ORG")
	urls, err := pkg.GetReposByOrg(org)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Unable to list repos for openfaas", http.StatusInternalServerError)
		return
	}

	bytesOut, _ := json.Marshal(urls)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesOut)
}
