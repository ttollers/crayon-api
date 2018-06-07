package routes

import (
	"net/http"
	"net/http/httptest"
	"github.com/julienschmidt/httprouter"
	"testing"
	"fmt"
)

func TestHandler(t *testing.T) {
	fmt.Print("HELLO");
	router := httprouter.New()
	router.GET("/book/:fromLang/:toLang/:bookName/:chapter", Handler)

	req, _ := http.NewRequest("GET", "/book/fr/en/harry-potter-1/chapter-1", nil)
	rr := httptest.NewRecorder();

	router.ServeHTTP(rr, req);

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}
