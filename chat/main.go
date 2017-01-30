package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/ehernandez-xk/bp_app/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/objx"
	"github.com/stretchr/signature"
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
	// new objetc to hold the variables that will be used in the template.
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	// now we are passing the data to inyect data in the template e.g. r.Host in {{.Host}} and {{.UserData.name}}
	t.templ.Execute(w, data)
}

// Added the Favicon
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func main() {
	var silent = flag.Bool("silent", false, "Silent the trancing logs")
	var addr = flag.String("addr", ":8080", "The address of the application")
	// parse the flags
	flag.Parse()

	// Environment variables
	ghClientid := os.Getenv("GITHUB_CLIENT_ID")
	ghClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	// setup gomniauth
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		github.New(ghClientid, ghClientSecret, "http://localhost:8080/auth/callback/github"),
	)

	r := newRoom()
	if *silent {
		fmt.Println("Tracing logs off")
		r.tracer = trace.Off()
	} else {
		fmt.Println("tracing logs active")
		r.tracer = trace.New(os.Stdout)
	}
	//http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
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
