package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando a API!")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":6001", r))
}