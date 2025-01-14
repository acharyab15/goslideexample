package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View

//START OMIT
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}

//END OMIT

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>OH NO!</h1> The URL you were trying to visit wasn't found!!")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	homeView = views.NewView("bootstrap", "views/home.gohtml")
// 	contactView = views.NewView("bootstrap", "views/contact.gohtml")
// 	faqView = views.NewView("faq", "views/faq.gohtml")

// 	r := mux.NewRouter()
// 	r.HandleFunc("/", home)
// 	r.HandleFunc("/contact", contact)
// 	r.HandleFunc("/faq", faq)

// 	var h http.Handler = http.HandlerFunc(notFound)
// 	r.NotFoundHandler = h
// 	http.ListenAndServe(":3000", r)
// }

//STARTMAIN OMIT
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)

}

//ENDMAIN OMIT
