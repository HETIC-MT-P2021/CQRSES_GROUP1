package controllers
import (
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Heyy :))")

}