package controllers

import (
	"admin-ambassador-checkout/src/database"
	"admin-ambassador-checkout/src/middlewares"
	"admin-ambassador-checkout/src/models"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	data["is_ambassador"] = strconv.FormatBool(strings.Contains(c.Path(), "/api/ambassador"))

	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	response, err := http.Post("http://172.17.0.1:8001/api/register", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}

	var user models.User

	json.NewDecoder(response.Body).Decode(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	isAmbassador := strings.Contains(c.Path(), "/api/ambassador")

	if isAmbassador {
		data["scope"] = "ambassador"
	} else {
		data["scope"] = "admin"
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	response, err := http.Post("http://172.17.0.1:8001/api/login", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}

	var res map[string]string

	json.NewDecoder(response.Body).Decode(&res)

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    res["jwt"],
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	req, err := http.NewRequest("GET", "http://172.17.0.1:8001/api/user", nil)

	if err != nil {
		return err
	}

	req.Header.Add("Cookie", "jwt="+c.Cookies("jwt", ""))

	client := http.Client{}

	response, err := client.Do(req)

	if err != nil {
		return err
	}

	var user models.User

	json.NewDecoder(response.Body).Decode(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, _ := middlewares.GetUserId(c)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	user.Id = id

	database.DB.Model(&user).Updates(&user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	id, _ := middlewares.GetUserId(c)

	user := models.User{}
	user.Id = id

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(&user)

	return c.JSON(user)
}
