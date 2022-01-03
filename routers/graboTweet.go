package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/josearlonmendez/twitter_arlon/bd"
	"github.com/josearlonmendez/twitter_arlon/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio u error al intentar grabar el tweet"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro el tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
