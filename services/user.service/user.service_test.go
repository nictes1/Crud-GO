package user_service_test

import (
	"testing"
	"time"

	m "github.com/nictes1/Golang-MongoDb-API/models"
	userService "github.com/nictes1/Golang-MongoDb-API/services/user.service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userId string

func TestCreate(t *testing.T) {

	oid := primitive.NewObjectID()
	userId = oid.Hex()

	user := m.User{
		ID:       oid,
		Name:     "Nicolas",
		Email:    "nikolastesone@gmail.com",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	err := userService.Create(user)

	if err != nil {
		t.Error("La prueba de persistencia del usuario a fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}

}

func TestRead(t *testing.T) {
	users, err := userService.Read()

	if err != nil {
		t.Error("Se ha presentado un error en la consulta de usuarios")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("La consulta no contiene datos")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name:  "Jesus Nicolas",
		Email: "nicolas@hotmail.com",
	}

	err := userService.Update(user, userId) //segundo parametro el obj id de la base de datos.

	if err != nil {
		t.Error("Error al tratar de actualizar el usuario")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito se actualizaron los datos.")
	}

}

func TestDelete(t *testing.T) {
	err := userService.Delete(userId)
	if err != nil {
		t.Error("Error al tratar de eliminar el usuario")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito se eliminaro los datos.")
	}

}
