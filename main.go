package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"github.com/ttollers/crayon-api/db"
)


func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fromLang := ps.ByName("fromLang")
	toLang := ps.ByName("toLang")
	bookName := ps.ByName("bookName")
	chapter := ps.ByName("chapter")

	objectPath := "books/" +  fromLang + "/" + toLang + "/" + bookName + "/" + chapter 

	err, data := db.GetObject(objectPath)


	if err != nil {		
		log.Print(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, data)
	}
}

func main(){
	router := httprouter.New()
	router.GET("/book/:fromLang/:toLang/:bookName/:chapter", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}


