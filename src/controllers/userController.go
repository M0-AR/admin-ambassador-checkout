package controllers

import (
	"admin-ambassador-checkout/src/database"
	"admin-ambassador-checkout/src/models"
	"admin-ambassador-checkout/src/services"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func Ambassador(c *fiber.Ctx) error {
	response, err := services.UserService.Get("users", c.Cookies("jwt", ""))

	if err != nil {
		return err
	}

	var users []models.User

	ambassadors := []models.User{}

	json.NewDecoder(response.Body).Decode(&users)

	for _, user := range users {
		if user.IsAmbassador {
			ambassadors = append(ambassadors, user)
		}
	}

	return c.JSON(ambassadors)
}

func Rankings(c *fiber.Ctx) error {
	rankings, err := database.Cache.ZRevRangeByScoreWithScores(context.Background(), "rankings", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	if err != nil {
		return err
	}

	result := make(map[string]float64)

	for _, ranking := range rankings {
		result[ranking.Member.(string)] = ranking.Score
	}

	return c.JSON(result)
}
