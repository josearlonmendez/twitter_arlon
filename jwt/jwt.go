package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/josearlonmendez/twitter_arlon/models"
)

func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("claveValidacionTwitor")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellido":         t.Apellido,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioWeb":         t.SitioWeb,
		"_ID":              t.ID.Hex(),
		"exp":              time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
