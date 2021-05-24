// Package goshopee realizes Shopee API v2
// https://open.shopee.com/documents?module=87&type=2&id=58&version=2
package goshopee

import (
	"net/http"
	"net/url"
)

const (
	UserAgent = "goshopee.v2/1.0.0"

	defaultApiPathPrefix = "api/v2"
	defaultApiVersion    = "v2"
	defaultHttpTimeout   = 10
)

// App represents basic app settings such as Api key, secret, scope, and redirect url.
// See oauth.go for OAuth related helper functions.
type App struct {
	PartnerID   int    `env:"SHOPEE_PARTNER_ID"`
	PartnerKey  string `env:"SHOPEE_PARTNER_KEY"`
	RedirectURL string `env:"SHOPEE_REDIRECT_URL"`
	APIURL      string `env:"SHOPEE_API_URL"`
	Client      *Client
}

type RateLimitInfo struct {
	RequestCount      int
	BucketSize        int
	RetryAfterSeconds float64
}

// Client manages communication with the Shopify API.
type Client struct {
	Client *http.Client
	log    LeveledLoggerInterface

	app App

	// Base URL for API requests.
	baseURL *url.URL

	// URL Prefix, defaults to "api" see WithVersion
	pathPrefix string

	// version you're currently using of the api, defaults to "v1"
	apiVersion string

	// max number of retries, defaults to 0 for no retries see WithRetry option
	retries  int
	attempts int

	RateLimits RateLimitInfo

	// Services used for communicating with the API
	Util UtilService
	Auth AuthService

	// Shop          ShopService
	// Item          ItemService
	// Variation     VariationService
	// ItemCategory  ItemCategoryService
	// ItemAttribute ItemAttributeService
	// Order         OrderService
	// Logistic      LogisticService
	// Discount DiscountService
}

// NewClient returns a new Shopify API client with an already authenticated shopname and
// token. The shopName parameter is the shop's myshopify domain,
// e.g. "theshop.myshopify.com", or simply "theshop"
// a.NewClient(shopName, token, opts) is equivalent to NewClient(a, shopName, token, opts)
func NewClient(app App, opts ...Option) *Client {
	baseURL, err := url.Parse(app.APIURL)
	if err != nil {
		panic(err)
	}

	c := &Client{
		Client:     &http.Client{},
		log:        &LeveledLogger{},
		app:        app,
		baseURL:    baseURL,
		apiVersion: defaultApiVersion,
		pathPrefix: defaultApiPathPrefix,
	}

	c.Util = &UtilServiceOp{client: c}
	c.Auth = &AuthServiceOp{client: c}
	
	// apply any options
	for _, opt := range opts {
		opt(c)
	}

	return c
}