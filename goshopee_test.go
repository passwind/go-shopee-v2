package goshopee

import (
	"fmt"
	"io/ioutil"

	"github.com/caarlos0/env"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	"github.com/prometheus/common/log"
)

const (
	maxRetries     = 3
	shopID = 1234567
	merchantID = 0
	accessToken = "accesstoken"
)

var (
	client *Client
	app    App
)

func setup() {
	err := godotenv.Load()
  if err != nil {
    log.Warn("Error loading .env file")
		app = App{
			PartnerID:      12345678,
			PartnerKey:   "hush",
			RedirectURL: "https://example.com/callback",
			APIURL: "https://partner.test-stable.shopeemobile.com",
		}
  }else {
		env.Parse(&app)
	}
	client = NewClient(app, 
		WithRetry(maxRetries))
	httpmock.ActivateNonDefault(client.Client)
}

func teardown() {
	httpmock.DeactivateAndReset()
}

func loadFixture(filename string) []byte {
	f, err := ioutil.ReadFile("fixtures/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}
	return f
}