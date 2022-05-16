package main

import (
	"log"
	"net/http"
	_ "url/internal/models"
	"url/internal/router"
)

func main() {

	log.Fatal(http.ListenAndServe(":8080", router.NewRouter()))
}
