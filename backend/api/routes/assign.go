package routes

import "github.com/gofiber/fiber/v2"

type request struct {
	URL    string          `json:"url"`
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

	//saving the ip in the database



	// checking all the responses
		// if not valid move ahead and print out a error in the frontend
		// if not move ahead
	// store the responses in the form of the map
	// check with the responses and pop out a button to generate it 
	// use the google api to generate in the google map and use it for furthur use


}
