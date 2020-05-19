package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sbdb-college/model"
	"strconv"
)

func getCollegeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	collegeId := r.URL.Query().Get("id")
	userId, _ := strconv.ParseUint(collegeId, 10, 64)
	college, err := model.Get(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(college)
	_, _ = w.Write(resp)
}

func postCollegeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var result model.College
	_ = json.Unmarshal(body, &result)
	result, err := model.Create(result)
	if err != nil {
		log.Println("Create college failed")
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		log.Println("College ", result.Name, "created")
	}
	response, err := json.Marshal(result)
	_, _ = w.Write(response)
}

func putCollegeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var result model.College
	_ = json.Unmarshal(body, &result)
	result, err := model.Put(result)
	if err != nil {
		log.Println("Update college failed")
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		log.Println("College ", result.Name, " updated")
	}
	response, err := json.Marshal(result)
	_, _ = w.Write(response)
}

func deleteCollegeHandler(w http.ResponseWriter, r *http.Request) {
	collegeId := r.URL.Query().Get("id")
	userId, _ := strconv.ParseUint(collegeId, 10, 64)
	err := model.Delete(userId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func CollegeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getCollegeHandler(w, r)
	case "POST":
		postCollegeHandler(w, r)
	case "PUT":
		putCollegeHandler(w, r)
	case "DELETE":
		deleteCollegeHandler(w, r)
	}
}
