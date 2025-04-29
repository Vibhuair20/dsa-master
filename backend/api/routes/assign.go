package routes

import (
	"fmt"

	"github.com/Vibhuair20/dsa-master/backend/api/database"
	"github.com/Vibhuair20/dsa-master/backend/api/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

type request struct {
	URL string 	`json:"url"`
	Email    string          `json:"email"`
	Topics map[string]bool 	 `json:"topics"`
	Dates  []string        	 `json:"dates"`
}

type response struct {
	Email 		string 				`json:"email"`
	Topics 		map[string]bool 	`json:"topics"`
	Dates 		[]string 			`json:"dates"`
	Sample 		map[string]bool 	`json:"sample"`
	URL string `json:"url"`
	// Generate 	string 				`json:"generate"`
}


func GenerateCalender(c *fiber.Ctx) error{
	body := new(request)

	if err := c.BodyParser(body), err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse the json"})
	}

	// check for the email and store it in the local redis
	
	// it retreives the email address from the dataset
	r2 := database.CreateClient(1)
	defer r2.Close()

	//getting the email address
	val, err := r2.Get(Database.Ctx, body.Email).Result()
	
	// first check the email is valid or not
	if !govalidator.IsEmail(body.Email){
		return c.Status(fiber.StatusBadRequest).JSIN(fiber.Mao{"error": "invalid email"})
	}

	if err != nil && err != redis.Nil{
		fmt.Println("error getting the value from the redis", err)
	}

	
	// if ip is not found in redis
	if err == redis.Nil{
		_ = r2.Set(database.Ctx, body.Email, body.Email, 0).Err()
		if err != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to store email in redis"})
		}
	}

	//! now checking the url for the question

	// check if the url provided is real or valid or not
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "the url provided here is not correct"})
	}
	// check for domain erroi

	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "not possoble"})
	}
	// enforce http , ssl
	body.URL = helpers.EnforceHTTP(body.URL)

}