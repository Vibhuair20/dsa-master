package routes

import "github.com/gofiber/fiber/v2"

type request struct {
	Email    string          `json:"email"`
	Topics map[string]bool `json:"topics"`
	Dates  []string        `json:"dates"`
}

type response struct {
	URL 		string 				`json:"url"`
	Topics 		map[string]bool 	`json:"topics"`
	Dates 		[]string 			`json:"dates"`
	Sample 		map[string]bool 	`json:"sample"`
	Generate 	string 				`json:"generate"`
}


func GenerateCalender(c *fiber.Ctx) error{
	body := new(request)

	// body request from the user to send the request
	if err := c.BodyParser(body), err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// selecting the options

	// applying the logic to the functions
	// if possible yes or not

	// then it predicts out the schedule for the revision and print it in the frontend


	// if yes then it triggers another function with the button
	// the button is enabled
	// after enabling the button it asks one more time with the details
	
	// when the button is click proceed

	// google calender with the api is integrated and added to the calender




}
