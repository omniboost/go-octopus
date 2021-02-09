package octopus_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	guestline "github.com/omniboost/go-octopus"
)

var (
	client *guestline.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	softwareHouse := os.Getenv("SOFTWARE_HOUSE")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	client = guestline.NewClient(nil, softwareHouse, username, password)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
