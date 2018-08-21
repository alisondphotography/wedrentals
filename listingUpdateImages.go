package main

import (
	"net/http"
	"html/template"
)

var listingUpdateImagesTemplate = template.Must(template.ParseFiles("template/root.html", "template/listingUpdateImages.html"))
func listingUpdateImagesHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	_, err := sessionIDCookie(r)
	if err == nil {
		p.IsLoggedIn = true
	}
	renderTemplate(w, listingUpdateImagesTemplate, p)
}
