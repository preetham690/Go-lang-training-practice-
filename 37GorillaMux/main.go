package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"httprequest/models"
	"net/http"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

var (
	PORT string // command line coding
)

// we are using gorilla mux over net/http because gorilla/mux offers better way of handing request funtions
func main() {
	flag.StringVar(&PORT, "port", "50080", "--port=50080 or -port=50080")
	flag.Parse()

	r := mux.NewRouter() //router

	srv := http.Server{
		Addr:         ":" + PORT,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	//to handle functions

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.HandleFunc("/Greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		fmt.Println(param["name"])
		fmt.Fprintln(w, param["name"]) // prints to the screen eg:postman
	})

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func Add(w http.ResponseWriter, r *http.Request) {
	//if it's not a post request
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Not Implemented"))
		return
	}

	file, err := os.ReadFile("employee/employee-id") //here file will have bytes

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	str := string(file)
	_id, err := strconv.Atoi(str)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(_id)

	e := new(models.Employee)

	err = json.NewDecoder(r.Body).Decode(e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = e.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	e.Id = _id + 1
	f, err = os.OpenFile("employee/"+strconv.Itoa(e.Id), syscall.O_RDWR, 0777)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

}
