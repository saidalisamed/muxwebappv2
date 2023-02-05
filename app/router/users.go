package router

import (
	"net/http"

	"github.com/saidalisamed/muxwebappv2/app/common"
)

func (ar *AppRouter) usersHandler(w http.ResponseWriter, r *http.Request) {
	users := ar.State.DB.GetUsers()
	common.RespondJSON(w, http.StatusOK, users)
}
