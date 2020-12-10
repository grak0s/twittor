package routers

import (
	"encoding/json"
	"net/http"

	"github.com/twittor/db"
	"github.com/twittor/jwt"
	"github.com/twittor/models"
)

//Login ruta que valida y genera el jwt
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "usuario y/o contraseña invalido"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El usuario de email es requerido", 400)
		return
	}

	documento, existe := db.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña invaldo", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Se ha generado un error al generar el token"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
