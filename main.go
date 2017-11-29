package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"time"
)

type IamInfo struct {
	Code    string
	LastUpdated time.Time
	InstanceProfileArn string
	InstanceProfileId string
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/latest/meta-data/iam/info", iamInfo).Methods("GET")
	rtr.HandleFunc("/latest/meta-data/iam/security-credentials/{name:[a-z]+}", profile).Methods("GET")

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

func iamInfo(w http.ResponseWriter, r *http.Request) {
	profile := IamInfo{"Success",
				time.Now(),
				"arn:aws:iam::008049814176:instance-profile/database",
				"AIPA00000000000000000"}

	js, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func profile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Profile " + name))
}