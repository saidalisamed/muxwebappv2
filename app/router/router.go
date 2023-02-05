package router

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/saidalisamed/muxwebappv2/app/state"
)

type AppRouter struct {
	State     *state.AppState
	Templates *template.Template
	Router    *mux.Router
}

func New(appState *state.AppState) *AppRouter {
	ar := &AppRouter{}
	ar.State = appState
	ar.Router = mux.NewRouter()
	ar.setupRoutes()
	ar.parseTemplates("app/templates/*")
	ar.configureStaticFiles("app/assets", "/static/")
	return ar
}

func (ar *AppRouter) setupRoutes() {
	ar.Router.HandleFunc("/", ar.homeHandler)
	ar.Router.HandleFunc("/users", ar.usersHandler)
	ar.Router.HandleFunc("/login", ar.loginHandler)
}

func (ar *AppRouter) parseTemplates(path string) {
	template := template.New("")
	_, err := template.ParseGlob(path)
	if err != nil {
		log.Println(err)
	}
	ar.Templates = template
}

func (ar *AppRouter) configureStaticFiles(dir string, path string) {
	fs := http.FileServer(http.Dir(dir))
	ar.Router.PathPrefix(path).Handler(http.StripPrefix(path, fs))
}
