package main

import (
	"net/http"
	"html/template"
)

var listingUpdateTemplate = template.Must(template.ParseFiles("template/root.html", "template/listingUpdate.html"))
func listingUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	_, err := sessionIDCookie(r)
	if err == nil {
		p.IsLoggedIn = true
	}
	renderTemplate(w, listingUpdateTemplate, p)
}
