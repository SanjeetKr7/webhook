package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var schoolDetail []SchoolDetail
var sd []StudentDetail
var id int

func server() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/addWebhookEvent", addWebhookEvent)
	http.HandleFunc("/addStudent", addStudent)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "welcome to Student Information !! ")
}

func addWebhookEvent(w http.ResponseWriter, r *http.Request) {
	var sc SchoolDetail
	byteValue, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("not able to read body := ", err)
	}
	json.Unmarshal(byteValue, &sc)
	uniqueId := isUnique(schoolDetail, sc.SchoolId)
	if uniqueId != false {
		schoolDetail = append(schoolDetail, sc)
	}
	byteVal, err := json.Marshal(schoolDetail)
	if err != nil {
		fmt.Println("not marshal := ", err)
	}
	io.WriteString(w, string(byteVal))
}

func isUnique(schoolDetail []SchoolDetail, id int) bool {
	for _, sd := range schoolDetail {
		if sd.SchoolId == id {
			return false
		}
	}
	return true
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var sc StudentDetail
	byteValue, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error")
	}

	json.Unmarshal(byteValue, &sc)
	if len(schoolDetail) == 0 {
		io.WriteString(w, "No college detail is added \n")
	}
	for _, s := range schoolDetail {
		if s.SchoolId == sc.SchoolId {
			// s.WebHook.EndPointUrl
			resp, err := http.Post(s.WebHook.EndPointUrl, "application/json", bytes.NewReader(byteValue))
			if err != nil {
				fmt.Println("not able to create webhook ", err)
			} else {
				b, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("not able to read return data ")
				}
				io.WriteString(w, fmt.Sprintf("The response is \n %s ! \n response code is %d ", string(b), resp.StatusCode))
			}

		} else {
			io.WriteString(w, fmt.Sprintf("The college id %d is not register ", sc.SchoolId))
		}
	}
	sd = append(sd, sc)
	byteVal, err := json.Marshal(sd)
	if err != nil {
		fmt.Println("not marshal ", err)
	}
	io.WriteString(w, string(byteVal))
}
