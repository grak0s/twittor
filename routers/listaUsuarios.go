package routers

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/grak0s/twittor/db"
)

//ListaUsuarios leo a lista de usuarios
func ListaUsuarios(w http.ResponseWriter, r*http.Request){
	typeUser	:= r.URL.Query().Get("type")
	page		:= r.URL.Query().Get("page")
	search		:= r.URL.Query().Get("search")

	pagTemp, err :=strconv.Atoi(page)

if err != nil{
	http.Error(w, "Enviar pagina mayor a cero", http.StatusBadRequest)
	return
}	
	pag := int64(pagTemp)

	result, status := db.LeoUsuariosTodos(IDUsuario, pag, search,typeUser)

if status == false {

	http.Error(w,"Error al leer los usuarios", http.StatusBadRequest)
	return

}

w.Header().Set("Content-Type", "appication/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(result)


}