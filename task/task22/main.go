/*The 68-http example is only rewriting the file. Please write it in a way that every request should be appened to the file.

Employee Id --> Should be auto incremented.

id no need to give manually in a http request. It has to be automatically taken.

delete a row based on id of an employee

sample query param : http://localhost:9091/employee/delete?id=1 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
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

type EmployeeHandler struct{}
type Employee struct {
	Emp_Num      string `json:"empnum"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobile       string `json:"mob"`
	Status       string `json:"status"`
	LastModified int    `json:"lastmodified"`
}

func EmpDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		e := new(Employee)

		err := json.NewDecoder(r.Body).Decode(&e) //takes data from request body
		//use json decoder and decode it to a greet struct variable
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if err := e.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		//e.Status = "active"
		//e.LastModified = int(time.Now().UTC().Unix())

		//here we are writing into the file

		if bytes, err := e.ToBytes(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		} else {
			// f, err := os.OpenFile("output.txt", syscall.O_RDWR, 0644)
			// if err != nil {
			// 	w.WriteHeader(http.StatusBadRequest)
			// 	w.Write([]byte(err.Error()))
			// 	return
			// }

			file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			}
			defer file.Close()

			if _, err = file.Write(bytes); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			}

			_, err = file.WriteString(string(bytes))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			} else {
				w.Write([]byte("Data successfully stored in the file"))
				w.WriteHeader(201)
				//w.WriteHeader(http.StatusCreated)
				return
			}
		}

		// fmt.Fprintln(w, "Data has been read", e)
	}
}

func (eh *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values.Get("id")
	if id == "" {
		w.Write([]byte("invalid id"))
		w.WriteHeader(400)
		return
	}
	//fmt.Println(id[0])
	fmt.Fprintln(w, "id query param is-> ", id)

	// Delete that row from the file
	// if no row with that id then it should say "details not found"

}

func (e *Employee) Validate() error {

	if e.Emp_Num == "" || e.Name == "" || e.Email == "" || e.Mobile == "" {
		return errors.New("invalid input data")
	}
	return nil
}

func (e *Employee) ToBytes() ([]byte, error) {
	return json.Marshal(e)
}
