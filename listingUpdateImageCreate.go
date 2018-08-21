package main

import (
	"net/http"
	"html/template"
)

var listingUpdateImageCreateTemplate = template.Must(template.ParseFiles("template/root.html", "template/listingUpdateImageCreate.html"))
func listingUpdateImageCreateHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	_, err := sessionIDCookie(r)
	if err == nil {
		p.IsLoggedIn = true
	}
	renderTemplate(w, listingUpdateImageCreateTemplate, p)
}
