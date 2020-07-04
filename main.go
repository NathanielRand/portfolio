package main

import (
	"fmt"
	"net/http"
	"log"
	"os"

	"github.com/NathanielRand/portfolio/views"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	homeView     *views.View
	contactView  *views.View
	socialsView  *views.View
	projectsView *views.View
	skillsView   *views.View
	resumeView   *views.View
	hobbiesView  *views.View
	travelsView  *views.View
)

type ContactForm struct {		
	Email     string `schema:"email"`
	Subject   string `schema:"subject"`
	Message   string `schema:"message"`
}

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

func skills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := skillsView.Template.ExecuteTemplate(w, skillsView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing skillsView template: %v\n", err)
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

func parseForm(r *http.Request, dst interface{}) error {
	// Parse the HTML form that was submitted
	// and store the data in the PostForm field of the http.Request.
	if err := r.ParseForm(); err != nil {
		return err
	}

	// Initialize our decoder.
	dec := schema.NewDecoder()	

	// Call the Decode method on our decode and
	// pass in the SignupForm as our destination and
	// r.PostForm as our data source. Check for errors.
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}

	// Return nil when no errors occur.
	return nil
}

func contactEmail(w http.ResponseWriter, r *http.Request) {	
	var form ContactForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	// form.Email
	// form.Subject
	// form.Message

	from := mail.NewEmail("Nathaniel Rand", "nathanieljrand@gmail.com")
	subject := form.Subject
	to := mail.NewEmail("Nathaniel Rand", "nathanieljrand@gmail.com")
	plainTextContent := form.Message
	htmlContent := "<strong>form.Message</strong>"
	
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	
	response, errSend := client.Send(message)
	if errSend != nil {
		log.Println(errSend)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		fmt.Println(form.Email)
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
	skillsView = views.NewView("materialize", "views/skills.gohtml")
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
	r.HandleFunc("/contact", contactEmail).Methods("POST")
	r.HandleFunc("/socials", socials)
	r.HandleFunc("/projects", projects)
	r.HandleFunc("/skills", skills)
	r.HandleFunc("/resume", resume)
	r.HandleFunc("/hobbies", hobbies)
	r.HandleFunc("/travels", travels)

	// Web Server
	http.ListenAndServe(":8080", r)
}
