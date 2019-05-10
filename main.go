package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/dennisg/movie-go/handlers"
)

var port = ":8080"

func init() {
	log.SetOutput(os.Stdout)

	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}
}


func main() {
	log.Fatal(http.ListenAndServe(port, nil))
}

