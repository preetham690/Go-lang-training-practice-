package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Employee struct {
	Id           String `json:"id`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobile       string `json:"mob"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastmodified"`
}

func (e *Employee) Validate() error {
	if e.Name == "" || e.Email == "" || e.Mobile == "" {
		fmt.Println(http.StatusNotAcceptable)
		return errors.New("invalid input data")
	}
	return nil
}

func (e *Employee) ToBytes() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Employee) ToString() (string, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
