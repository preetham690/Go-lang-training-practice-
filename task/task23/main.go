package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/Get", GetDetails)

	http.HandleFunc("/Delete", DeleteData)
}

func GetDetails(w http.ResponseWriter, r *http.Request) {

}

func DeleteData(w http.ResponseWriter, r *http.Request) {

}

type Employee struct {
	Emp_Num      string `json:empnum`
	Name         string `json:name`
	Email        string `json:email`
	Mobile       string `json:mob`
	Status       string `json:status`
	LastModified string `json:lastmodified`
}
