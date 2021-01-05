package controllers
import (
	"net/http"

	"CQRSES/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Heyy :))")

}
