package routers

import (
	"encoding/json"
	"net/http"

	"github.com/josearlonmendez/twitter_arlon/bd"
	"github.com/josearlonmendez/twitter_arlon/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	status, err := bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio u error al intentar grabar el registro"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
