package services

import "github.com/M0-AR/go-user-service"

var UserService services.Service

func Setup() {
	UserService = services.CreateService("http://users-ms:8000/api/")
}
