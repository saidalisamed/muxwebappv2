package state

import (
	"github.com/gorilla/sessions"
	"github.com/saidalisamed/muxwebappv2/app/config"
	"github.com/saidalisamed/muxwebappv2/app/db"
)

type AppState struct {
	DB       *db.Manager
	Sessions *sessions.CookieStore
	Config   *config.Configuration
}
