package main

import (
	"admin-ambassador-checkout/src/database"
	"admin-ambassador-checkout/src/models"
	"admin-ambassador-checkout/src/services"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

func main() {
	database.Connect()
	database.SetupRedis()

	ctx := context.Background()

	response, err := services.UserService.Get("users", "")

	if err != nil {
		panic(err)
	}

	var users []models.User

	json.NewDecoder(response.Body).Decode(&users)

	for _, user := range users {
		if user.IsAmbassador {
			ambassador := models.Ambassador(user)
			ambassador.CalculateRevenue(database.DB)

			database.Cache.ZAdd(ctx, "rankings", &redis.Z{
				Score:  *ambassador.Revenue,
				Member: user.Name(),
			})
		}
	}
}
