package db

import (
	"github.com/grak0s/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin reqliza el chequeo en la BD
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}

	return usu, true

}
