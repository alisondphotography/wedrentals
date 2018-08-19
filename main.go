package main

import (
	"net/http"
	"html/template"
	"log"
	"errors"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", makeAuthHandler("accountInsert.html"))
	http.HandleFunc("/login", makeAuthHandler("accountLogIn.html"))
	http.HandleFunc("/login/forgot-password", makeAuthHandler("logInForgotPassword.html"))
	http.HandleFunc("/logout", logOutHandler)
	http.HandleFunc("/account", makeHandler("account.html"))
	http.HandleFunc("/account/update/email", makeHandler("accountUpdateEmail.html"))
	http.HandleFunc("/account/update/password", makeHandler("accountUpdatePassword.html"))
	http.HandleFunc("/vendor/profile/create", makeHandler("profileInsert.html"))
	http.HandleFunc("/vendor/profile/update", makeHandler("profileUpdate.html"))
	http.HandleFunc("/vendor/post", makeHandler("listingInsert.html"))
	http.HandleFunc("/vendor/post/update", makeHandler("listingUpdate.html"))
	http.HandleFunc("/vendor/listings", makeHandler("vendorListings.html"))
	http.HandleFunc("/listings", makeListingsHandler("listings.html"))
	http.HandleFunc("/listing", makeHandler("listing.html"))
	http.HandleFunc("/profile", makeHandler("profile.html"))
	http.HandleFunc("/account/faves", makeHandler("faves.html"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Page struct {
	IsLoggedIn bool
}

var indexTemplate = template.Must(template.ParseFiles("template/index.html"))
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	var p Page
	_, err := sessionIDCookie(r)
	if err == nil {
		p.IsLoggedIn = true
	}
	renderTemplate(w, indexTemplate, p)
}

func makeAuthHandler(f string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			http.SetCookie(w, &http.Cookie{Name: "session_id", Value: "12345", Path: "/"})
			http.Redirect(w, r, "/account/faves", http.StatusFound)
			return
		}

		t, err := template.ParseFiles("template/root.html", "template/"+f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var p Page
		_, err = sessionIDCookie(r)
		if err == nil {
			p.IsLoggedIn = true
		}
		renderTemplate(w, t, p)
	}
}

func makeListingsHandler(f string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("template/root.html", "template/"+f, "template/listings-section.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var p Page
		_, err = sessionIDCookie(r)
		if err == nil {
			p.IsLoggedIn = true
		}
		renderTemplate(w, t, p)
	}
}

func makeHandler(f string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("template/root.html", "template/"+f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var p Page
		_, err = sessionIDCookie(r)
		if err == nil {
			p.IsLoggedIn = true
		}
		renderTemplate(w, t, p)
	}
}

func logOutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session_id", Value: "", Path: "/", Expires: time.Now().Add(time.Hour * -1)})
	http.Redirect(w, r, "/login", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, t *template.Template, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sessionIDCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", err
	}
	if cookie.Value == "" {
		return "", errors.New("session_id cookie is empty")
	}
	return cookie.Value, nil
}