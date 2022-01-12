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

	baseURLString := os.Getenv("OCTO_BASE_URL")
	softwareHouse := os.Getenv("OCTO_SOFTWARE_HOUSE")
	username := os.Getenv("OCTO_USERNAME")
	password := os.Getenv("OCTO_PASSWORD")
	debug := os.Getenv("OCTO_DEBUG")
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
