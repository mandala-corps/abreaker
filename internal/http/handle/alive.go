package handle

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Alive(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	writeJson(w, &R{
		"alive": "i think that yes",
	})
}
