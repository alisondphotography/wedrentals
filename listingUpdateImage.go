package main

import (
	"net/http"
	"html/template"
)

var listingUpdateImageTemplate = template.Must(template.ParseFiles("template/root.html", "template/listingUpdateImage.html"))
func listingUpdateImageHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	_, err := sessionIDCookie(r)
	if err == nil {
		p.IsLoggedIn = true
	}
	renderTemplate(w, listingUpdateImageTemplate, p)
}
