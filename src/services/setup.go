package services

var UserService Service

func Setup() {
	UserService = CreateService("http://172.17.0.1:8001/api/")
}
