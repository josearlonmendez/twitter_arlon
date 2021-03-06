package routers

import (
	"encoding/json"
	"net/http"

	"github.com/josearlonmendez/twitter_arlon/bd"
	"github.com/josearlonmendez/twitter_arlon/models"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion
	status, err := bd.ConsultoRelacion(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
