package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	PORT string
)

func init() {
	logger := log.New(os.Stdout, "", log.Lshortfile|log.Ldate|log.Ltime)
}

func main() {

	flag.StringVar(&PORT, "port", "50080", "--port=50080 or -port=50080")
	flag.Parse()

	router := mux.NewRouter()

	srv := http.Server{
		Addr:    ":" + PORT,
		Handler: router,
	}

	logger.Println("Server started and liostening", PORT)

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalln(err)
	}
}
