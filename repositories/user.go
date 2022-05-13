package repositories

import (
	"assignment1/models"
	"assignment1/params"
	"errors"
	"time"
)

func InsertUser(req *params.CreateUser) models.User {
	newUser := models.User{
		Nama:      req.Nama,
		Email:     req.Email,
		Pekerjaan: req.Pekerjaan,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return newUser
}

func FindUserByNomorAbsen(users []models.User, absen int) (models.User, error) {
	currentUser := models.User{}
	if (len(users) < absen) || (absen <= 0) {
		return currentUser, errors.New("user tidak ditemukan")
	}

	currentUser = users[absen-1]

	return currentUser, nil
}
