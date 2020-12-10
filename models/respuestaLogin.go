package models

//RespuestaLogin tiene el token que dvuelve el login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
