package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/sessions"
	"github.com/saidalisamed/muxwebappv2/app/config"
	"github.com/saidalisamed/muxwebappv2/app/db"
	"github.com/saidalisamed/muxwebappv2/app/router"
	"github.com/saidalisamed/muxwebappv2/app/state"
)

var appConfig = config.ReadConfig()

var appState = &state.AppState{
	DB:       db.NewManager(appConfig.DBUsername, appConfig.DBPassword, appConfig.DBHost, appConfig.DBName, appConfig.DBPort),
	Sessions: sessions.NewCookieStore([]byte(appConfig.AppSecret)),
	Config:   appConfig,
}

var appRouter = router.NewRouter(appState)

func Run() {
	loggedRouter := handlers.LoggingHandler(os.Stdout, appRouter.Router)
	addr := fmt.Sprintf("%s:%d", appState.Config.AppListenIP, appState.Config.AppListenPort)

	log.Printf("Starting web server on http://%s\n", addr)
	err := http.ListenAndServe(addr, handlers.RecoveryHandler()(loggedRouter))

	if err != nil {
		log.Println(err)
	}
}
