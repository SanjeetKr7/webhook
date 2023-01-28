package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var g []StudentDetail

func Server() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/webhook/studentAdded", studentAdded)
	log.Fatal(http.ListenAndServe(":9091", nil))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "welcome abc school!!!!!!!!!!")
}
func studentAdded(w http.ResponseWriter, r *http.Request) {
	byteValue, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error")
	}
	var sc StudentDetail
	json.Unmarshal(byteValue, &sc)
	g = append(g, sc)
	byteVal, err := json.Marshal(g)
	if err != nil {
		fmt.Println("not marshal")
	}
	fmt.Println("All the student added is ", string(byteVal))
	io.WriteString(w, fmt.Sprintf("webhook data received %s ! ", string(byteValue)))
}
