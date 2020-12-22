package controller

import (
	"github.com/thorstenkloehn/lernen/Server/ahrensburg/model"
	"net/http"
	"text/template"
)

var view, _ = template.ParseGlob("ahrensburg/views/*")
var website = model.Website{}

func Startseite(w http.ResponseWriter, r *http.Request) {

	website.Titel = "ahrensburg.digital"

	view.ExecuteTemplate(w, "index.html", website)

}
