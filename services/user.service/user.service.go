package user_service

import (
	m "github.com/nictes1/Golang-MongoDb-API/models"
	userRepository "github.com/nictes1/Golang-MongoDb-API/repositories/user_repository"
)

func Create(user m.User) error {

	err := userRepository.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Users, error) {

	users, err := userRepository.Read()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func Update(user m.User, userId string) error {
	err := userRepository.Update(user, userId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {
	err := userRepository.Delete(userId)

	if err != nil {
		return err
	}

	return nil
}
