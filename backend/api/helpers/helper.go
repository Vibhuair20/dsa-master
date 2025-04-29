package helpers

import (
	"context"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// EnforceHTTP ensures URLs start with http://
func EnforceHTTP(url string) string {
	if !strings.HasPrefix(url, "http") {
		return "http://" + url
	}
	return url
}

// RemoveDomainError checks if URL is same as configured domain
func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}
	newURL := strings.Replace(url, "https://", "", 1)
	newURL = strings.Replace(newURL, "http://", "", 1) // fixed: should remove http:// not https:// again
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]
	return newURL != os.Getenv("DOMAIN")
}

// GetGoogleOAuthConfig loads client credentials from credentials.json
func GetGoogleOAuthConfig() *oauth2.Config {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		panic("Unable to read credentials.json")
	}
	conf, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		panic("Unable to parse credentials.json")
	}
	return conf
}

// GetUserEmailFromToken retrieves the user's primary calendar email
func GetUserEmailFromToken(accessToken string) (string, error) {
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return "", err
	}
	cals, err := srv.CalendarList.List().Do()
	if err != nil || len(cals.Items) == 0 {
		return "", err
	}
	return cals.Items[0].Id, nil
}
