package main

import (
	"fmt"
	"net/http"

	"github.com/NathanielRand/portfolio/views"

	"github.com/gorilla/mux"
)

var (
	homeView *views.View
	contactView *views.View
	socialsView *views.View
	projectsView *views.View
	educationView *views.View
	resumeView *views.View
	hobbiesView *views.View
	travelsView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing homeView template: %v\n", err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func socials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := socialsView.Template.ExecuteTemplate(w, socialsView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing socialsView template: %v\n", err)
	}
}

func projects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := projectsView.Template.ExecuteTemplate(w, projectsView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing projectsView template: %v\n", err)
	}
}

func education(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := educationView.Template.ExecuteTemplate(w, educationView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing educationView template: %v\n", err)
	}
}

func resume(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := resumeView.Template.ExecuteTemplate(w, resumeView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing resumeView template: %v\n", err)
	}
}

func hobbies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := hobbiesView.Template.ExecuteTemplate(w, hobbiesView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing hobbiesView template: %v\n", err)
	}
}

func travels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := travelsView.Template.ExecuteTemplate(w, travelsView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing travelsView template: %v\n", err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/favicon.ico")
}

func main() {
	// Call NewView func from views package
	homeView = views.NewView("materialize", "views/home.gohtml")
	contactView = views.NewView("materialize", "views/contact.gohtml")
	socialsView = views.NewView("materialize", "views/socials.gohtml")
	projectsView = views.NewView("materialize", "views/projects.gohtml")
	educationView = views.NewView("materialize", "views/education.gohtml")
	resumeView = views.NewView("materialize", "views/resume.gohtml")
	hobbiesView = views.NewView("materialize", "views/hobbies.gohtml")
	travelsView = views.NewView("materialize", "views/travels.gohtml")
	
	// Create router
	r := mux.NewRouter()

	// Favicon
	r.HandleFunc("/favicon.ico", faviconHandler)

	// Assest Routes
	assetHandler := http.FileServer(http.Dir("./assets/"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	// Define paths with assigned functions
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/socials", socials)
	r.HandleFunc("/projects", projects)
	r.HandleFunc("/education", education)
	r.HandleFunc("/resume", resume)
	r.HandleFunc("/hobbies", hobbies)
	r.HandleFunc("/travels", travels)

	// Web Server
	http.ListenAndServe(":8080", r)
}