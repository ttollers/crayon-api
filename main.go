package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"github.com/ttollers/crayon-api/routes"
)

func main(){
	router := httprouter.New()
	router.GET("/book/:fromLang/:toLang/:bookName/:chapter", routes.Handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}


