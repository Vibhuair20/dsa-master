package auth

import (
	"context"
	"encoding/json"

	"github.com/Vibhuair20/dsa-master/backend/api/database"
	"github.com/Vibhuair20/dsa-master/backend/api/helpers"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"google.golang.org/api/oauth2/v2"
)

// gets the oauth2 configuration  in helpers
// generates the auth token with a atate(to prevent crssf)
// redirects to google oauth screen
func GoogleLogin(c *fiber.Ctx) error {
	conf := helpers.GetGoogleOAuthConfig()
	url := conf.AuthCodeURL("state-taken", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

// google callback
func GoogleCallBack(c *fiber.Ctx) error {
	code := c.Query("code")
	conf := helpers.GetGoogleOAuthConfig()

	// sxchange the code for tokens

	tok, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(500).SendString("Token exchange failed")
	}

	email, err := helpers.GetUserEmailFromToken(tok.AccessToken)
	if err != nil {
		return c.Status(500).SendString("failed to fetch user email")
	}

	rdb := database.CreateClient(1)
	defer rdb.Close()

	tokenJson, _ := json.Marshal(tok)
	err = rdb.Set(database.Ctx, email, tokenJson, 0).Err()
	if err != nil {
		return c.Status(500).SendString("failed to store the token")
	}
	return c.JSON(fiber.Map{"message": "Login success", "email": email})
	// store the token in the redis with the email as the key

}
