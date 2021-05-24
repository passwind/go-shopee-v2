package goshopee

import (
	"github.com/caarlos0/env"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	"github.com/prometheus/common/log"
)

const (
	testApiVersion = "9999-99"
	maxRetries     = 3
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
		WithVersion(testApiVersion),
		WithRetry(maxRetries))
	httpmock.ActivateNonDefault(client.Client)
}

func teardown() {
	httpmock.DeactivateAndReset()
}