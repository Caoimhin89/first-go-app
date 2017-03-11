package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/Caoimhin89/first-go-app/common"
	"github.com/Caoimhin89/first-go-app/routers"
)

func main() {
	common.StartUp()
	r := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(r)

	log.Println("Listening...")
	http.ListenAndServe(":8080", n)
}