package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Erick-Alen/hexagonal-architecture/adapters/web/handler"
	"github.com/Erick-Alen/hexagonal-architecture/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {

	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())
	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)
	// code to run the web server
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
