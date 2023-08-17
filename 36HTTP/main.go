package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// the safest port numbers are from 30000-65000
func main() {
	fmt.Println("Server started and listing in port 9091")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	http.HandleFunc("/vs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is Victoria Secret")
	})

	http.HandleFunc("/GreetBy", GreetBy)

	http.HandleFunc("/Greet", SayHi)

	go http.ListenAndServe(":9092", nil)

	err := http.ListenAndServe("0.0.0.0:9091", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(err, "hello world")
}

func SayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Good afternoon")
}

func GreetBy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		g := Greet{}

		err := json.NewDecoder(r.Body).Decode(&g) //takes data from request body
		//use json decoder and decode it to a greet struct variable
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
		}
		fmt.Fprintln(w, "Data has been read", g)
	} else {
		w.Write([]byte("Unimplemeted method"))
		w.WriteHeader(http.StatusNotImplemented)
	}
}

type Greet struct {
	Message string `json : "message"` // serialization or deserialization
	Name    string `json : "name"`
}

type Employee struct {
	Emp_Num      string `json:empnum`
	Name         string `json:name`
	Email        string `json:email`
	Mobile       string `json:mob`
	Status       string `json:status`
	LastModified string `json:lastmodified`
}
