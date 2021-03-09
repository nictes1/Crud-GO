package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User datos del usuario
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json: "id,omitempty"`
	Name     string             `json: "name"`
	Email    string             `json: "email"`
	CreateAt time.Time          `bson:"create_at" json: "create_at"`
	UpdateAt time.Time          `bson:"uptdate_up" json: "update_at, omitempty"`
}

//lista de usuarios
type Users []*User
