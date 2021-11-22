package middlew

import (
	"net/http"

	"github.com/josearlonmendez/twitter_arlon/bd"
)

func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnections() == 0 {
			http.Error(w, "Conexion perdida con la BD ", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
