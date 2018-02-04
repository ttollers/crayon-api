package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"github.com/julienschmidt/httprouter"
	"log"
)


func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	BOOK_PATH := os.Getenv("CHAPTER_DIR");

	bookName := ps.ByName("bookName")
	nativeLang := ps.ByName("nativeLang")
	chapter := ps.ByName("chapter")
	fileName := ps.ByName("type") + ".json"

	filePath := BOOK_PATH + "/" +  bookName + "/" + nativeLang + "/" + chapter + "/" + fileName

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(data))
}

func main(){
	router := httprouter.New()
	router.GET("/book/:bookName/:nativeLang/:chapter/:type", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}


