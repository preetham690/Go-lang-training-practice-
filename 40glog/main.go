package main

import (
	"os"

	"github.com/golang/glog"
)

func main() {
	file, err := os.OpenFile("myLOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		glog.Fatal(err)
	}

	glog.CopyStandardLogTo(file.Name())

	glog.Infoln("Welcome to victoria Secret & co")
	glog.Error("No error but error log")
}
