package main

import (
	"fmt"
	"lakuinv5dash/models"
	"lakuinv5dash/routes"
	"lakuinv5dash/utils"
	"log"
	"net/http"
)

const (
	PORT = ":8184"
)

func main() {
	models.Connection()
	fmt.Printf("Listerning Port %s\n", PORT)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))

}
