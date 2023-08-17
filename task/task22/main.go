package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	fmt.Println("Server started and listing in port 9091")

	http.HandleFunc("/EmpDet", EmpDetails)

	err := http.ListenAndServe("0.0.0.0:9091", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

type Employee struct {
	Emp_Num      string `json:empnum`
	Name         string `json:name`
	Email        string `json:email`
	Mobile       string `json:mob`
	Status       string `json:status`
	LastModified string `json:lastmodified`
}

// Read implements io.Reader.
// func (Employee) Read(p []byte) (n int, err error) {
// 	panic("unimplemented")
// }

func EmpDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		e := new(Employee)

		err := json.NewDecoder(r.Body).Decode(&e) //takes data from request body
		//use json decoder and decode it to a greet struct variable
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
		}

		e.Status = "active"
		e.LastModified = time.Now().UTC().Unix()

		//here we are writing into the file

		// body, err := io.ReadAll(r.Body)
		// fmt.Println(body)
		// if err != nil {
		// 	log.Fatal("Can not write to the file")
		// }
		// if err := os.Write("output.txt", body, 0644); err != nil {
		// 	fmt.Println("Error", err)
		// }

		// f, err := os.Create("Output.txt")
		// if err != nil {
		// 	panic(err)
		// }
		// f.WriteString(g.Emp_Num)

		// fmt.Println(f)

		// f, err := os.Create("data.txt")

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// defer f.Close()

		// val := r.Body
		// data := (val)

		// _, err2 := f.Write(data)
		// if err2 != nil {
		// 	log.Fatal(err2)
		// }

		fmt.Fprintln(w, "Data has been read", e)
	} else {
		w.Write([]byte("Unimplemeted method"))
		w.WriteHeader(http.StatusNotImplemented)
	}
}
