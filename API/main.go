package main

import (
	router "API/src/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("iniciando APi")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":3030", r))
}
