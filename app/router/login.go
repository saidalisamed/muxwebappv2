package router

import (
	"fmt"
	"net/http"
)

func (ar *AppRouter) loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := ar.State.Sessions.Get(r, "session")

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		// password := r.FormValue("password")

		// Perform authentication here (e.g. check against database)

		session.Values["username"] = username
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, `
		<form action="/login" method="post">
			<input type="text" name="username">
			<input type="password" name="password">
			<input type="submit" value="Login">
		</form>
	`)
}
