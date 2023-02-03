package router

import (
	"net/http"
)

// HomeHandler is a function that serves the "/" route.
func (ar *AppRouter) homeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": ar.State.Config.AppTitle,
		"body":  "Welcome to my website",
		"users": ar.State.DB.GetUsers(),
	}

	err := ar.Templates.ExecuteTemplate(w, "home", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
