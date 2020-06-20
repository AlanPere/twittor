package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omintempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omintempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omintempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"FfechaNacimiento,omintempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioweb" json:"sitioWeb,omitempty"`
}
