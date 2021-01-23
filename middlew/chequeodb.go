package middlew

import (
	"net/http"

	"github.com/grak0s/twittor/db"
)

//ChequeoDB valida la conexion
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida a la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
