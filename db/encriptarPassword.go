package db

import "golang.org/x/crypto/bcrypt"

//EncriptarPassword encripta la clave del usuario en 8 capas
func EncriptarPassword(pass string) (string, error) {

	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err

}
