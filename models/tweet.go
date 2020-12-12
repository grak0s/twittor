package models

//Tweet captura el body del mensaje que nos llega
type Tweet struct {
	Mensaje string `bson:"Mensaje" json:"Mensaje"`
}
