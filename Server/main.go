package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/thorstenkloehn/lernen/Server/ahrensburg/controller"

	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

func main() {

	var docs, static, lernen, lernportal string

	flag.StringVar(&docs, "docs", "docs/book", "")
	flag.Parse()
	flag.StringVar(&static, "static", "static/", "")
	flag.Parse()
	flag.StringVar(&lernen, "lernen", "lernen/book/", "")
	flag.Parse()
	flag.StringVar(&lernportal, "lernportal", "lernportal/book/", "")
	flag.Parse()
	server := mux.NewRouter()
	server.HandleFunc("/", controller.Startseite).Host("ahrensburg.digital")
	server.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(docs)))).Host("webprogrammierung.org")
	server.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(static)))).Host("static.webprogrammierung.org")
	server.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(lernen)))).Host("lernen.webprogrammierung.org")
	server.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(lernportal)))).Host("deutsch.webprogrammierung.org")
	go http.ListenAndServe(":80", handlers.CompressHandler(server))

	http.Serve(autocert.NewListener("webprogrammierung.org", "static.webprogrammierung.org", "ahrensburg.digital", "lernen.webprogrammierung.org", "deutsch.webprogrammierung.org"), handlers.CompressHandler(server))
}
