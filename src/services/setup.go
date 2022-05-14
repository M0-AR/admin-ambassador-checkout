package services

var UserService Service

func Setup() {
	UserService = CreateService("http://users-ms:8000/api/")
}
