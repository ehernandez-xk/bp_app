package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServerHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	// Passing the request to inyect data in the template e.g. r.Host in {{.Host}}
	t.templ.Execute(w, r)
}

// Added the Favicon
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func main() {
	var addr = flag.String("addr", ":8080", "The address of the application")
	// parse the flags
	flag.Parse()

	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// get the room going
	go r.run()

	http.HandleFunc("/favicon.ico", faviconHandler)

	// start the web server
	log.Println("Starting the web server on:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
