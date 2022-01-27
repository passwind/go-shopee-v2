// Package goshopee realizes Shopee API v2
// https://open.shopee.com/documents?module=87&type=2&id=58&version=2
package goshopee

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	UserAgent = "goshopee.v2/1.0.0"

	defaultHttpTimeout = 10
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

	// max number of retries, defaults to 0 for no retries see WithRetry option
	retries  int
	attempts int

	RateLimits RateLimitInfo

	ShopID      uint64
	MerchantID  uint64
	AccessToken string

	// Services used for communicating with the API
	Util      UtilService
	Auth      AuthService
	Media     MediaSpaceService
	Product   ProductService
	Logistics LogisticsService
	Shop      ShopService
	Discount  DiscountService
	Order     OrderService
	Merchant  MerchantService
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
		Client:  &http.Client{},
		log:     &LeveledLogger{},
		app:     app,
		baseURL: baseURL,
	}

	c.Util = &UtilServiceOp{client: c}
	c.Auth = &AuthServiceOp{client: c}
	c.Media = &MediaSpaceServiceOp{client: c}
	c.Product = &ProductServiceOp{client: c}
	c.Logistics = &LogisticsServiceOp{client: c}
	c.Shop = &ShopServiceOp{client: c}
	c.Discount = &DiscountServiceOp{client: c}
	c.Order = &OrderServiceOp{client: c}
	c.Merchant = &MerchantServiceOp{client: c}

	// apply any options
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// A general response error that follows a similar layout to Shopify's response
// errors, i.e. either a single message or a list of messages.
type ResponseError struct {
	Status  int
	Message string
	Errors  []string
}

// GetStatus returns http  response status
func (e ResponseError) GetStatus() int {
	return e.Status
}

// GetMessage returns response error message
func (e ResponseError) GetMessage() string {
	return e.Message
}

// GetErrors returns response errors list
func (e ResponseError) GetErrors() []string {
	return e.Errors
}

func (e ResponseError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	sort.Strings(e.Errors)
	s := strings.Join(e.Errors, ", ")

	if s != "" {
		return s
	}

	return "Unknown Error"
}

// ResponseDecodingError occurs when the response body from Shopify could
// not be parsed.
type ResponseDecodingError struct {
	Body    []byte
	Message string
	Status  int
}

func (e ResponseDecodingError) Error() string {
	return e.Message
}

// An error specific to a rate-limiting response. Embeds the ResponseError to
// allow consumers to handle it the same was a normal ResponseError.
type RateLimitError struct {
	ResponseError
	RetryAfter int
}

// Creates an API request. A relative URL can be provided in urlStr, which will
// be resolved to the BaseURL of the Client. Relative URLS should always be
// specified without a preceding slash. If specified, the value pointed to by
// body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, relPath string, body, options, headers interface{}) (*http.Request, error) {
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	// Add custom options
	if options != nil {
		optionsQuery, err := query.Values(options)
		if err != nil {
			return nil, err
		}

		for k, values := range u.Query() {
			for _, v := range values {
				optionsQuery.Add(k, v)
			}
		}
		u.RawQuery = optionsQuery.Encode()
	}

	// A bit of JSON ceremony
	var js []byte = nil
	if body != nil {
		js, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)

	c.makeSignature(req)

	return req, nil
}

func (c *Client) WithShop(sid uint64, tok string) *Client {
	c.ShopID = sid
	c.AccessToken = tok
	return c
}

func (c *Client) WithMerchant(mid uint64, tok string) *Client {
	c.MerchantID = mid
	c.AccessToken = tok
	return c
}

func (c *Client) WithToken(tok string) *Client {
	c.AccessToken = tok
	return c
}

// https://open.shopee.com/documents?module=87&type=2&id=58&version=2
func (c *Client) makeSignature(req *http.Request) (string, int64) {
	ts := time.Now().Unix()
	path := req.URL.Path

	var baseStr string

	u := req.URL

	query := u.Query()
	query.Add("partner_id", fmt.Sprintf("%v", c.app.PartnerID))

	if c.ShopID != 0 {
		// Shop APIs: partner_id, api path, timestamp, access_token, shop_id
		baseStr = fmt.Sprintf("%d%s%d%s%d", c.app.PartnerID, path, ts, c.AccessToken, c.ShopID)
		query.Add("shop_id", fmt.Sprintf("%v", c.ShopID))
		query.Add("access_token", c.AccessToken)
	} else if c.MerchantID != 0 {
		// Merchant APIs: partner_id, api path, timestamp, access_token, merchant_id
		baseStr = fmt.Sprintf("%d%s%d%s%d", c.app.PartnerID, path, ts, c.AccessToken, c.MerchantID)
		query.Add("merchant_id", fmt.Sprintf("%v", c.MerchantID))
		query.Add("access_token", c.AccessToken)
	} else {
		// Public APIs: partner_id, api path, timestamp
		baseStr = fmt.Sprintf("%d%s%d", c.app.PartnerID, path, ts)
	}
	h := hmac.New(sha256.New, []byte(c.app.PartnerKey))
	h.Write([]byte(baseStr))
	result := hex.EncodeToString(h.Sum(nil))

	query.Add("timestamp", fmt.Sprintf("%v", ts))
	query.Add("sign", result)

	u.RawQuery = query.Encode()
	req.URL = u

	return result, ts
}

// doGetHeaders executes a request, decoding the response into `v` and also returns any response headers.
func (c *Client) doGetHeaders(req *http.Request, v interface{}, skipBody bool) (http.Header, error) {
	var resp *http.Response
	var err error

	retries := c.retries
	c.attempts = 0
	c.logRequest(req, skipBody)

	for {
		c.attempts++

		resp, err = c.Client.Do(req)
		c.logResponse(resp)
		if err != nil {
			return nil, err //http client errors, not api responses
		}

		respErr := CheckResponseError(resp)
		if respErr == nil {
			break // no errors, break out of the retry loop
		}

		// retry scenario, close resp and any continue will retry
		resp.Body.Close()

		if retries <= 1 {
			return nil, respErr
		}

		if rateLimitErr, isRetryErr := respErr.(RateLimitError); isRetryErr {
			// back off and retry

			wait := time.Duration(rateLimitErr.RetryAfter) * time.Second
			c.log.Debugf("rate limited waiting %s", wait.String())
			time.Sleep(wait)
			retries--
			continue
		}

		var doRetry bool
		switch resp.StatusCode {
		case http.StatusServiceUnavailable:
			c.log.Debugf("service unavailable, retrying")
			doRetry = true
			retries--
		}

		if doRetry {
			continue
		}

		// no retry attempts, just return the err
		return nil, respErr
	}

	c.logResponse(resp)
	defer resp.Body.Close()

	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&v)
		if err != nil {
			return nil, err
		}
	}

	return resp.Header, nil
}

// skipBody: if upload image, skip log its binary
func (c *Client) logRequest(req *http.Request, skipBody bool) {
	if req == nil {
		return
	}
	if req.URL != nil {
		c.log.Debugf("%s: %s", req.Method, req.URL.String())
	}
	if !skipBody {
		c.logBody(&req.Body, "SENT: %s")
	}
}

func (c *Client) logResponse(res *http.Response) {
	if res == nil {
		return
	}
	c.log.Debugf("RECV %d: %s", res.StatusCode, res.Status)
	c.logBody(&res.Body, "RESP: %s")
}

func (c *Client) logBody(body *io.ReadCloser, format string) {
	if body == nil || *body == nil {
		return
	}
	b, _ := ioutil.ReadAll(*body)
	if len(b) > 0 {
		c.log.Debugf(format, string(b))
	}
	*body = ioutil.NopCloser(bytes.NewBuffer(b))
}

func wrapSpecificError(r *http.Response, err ResponseError) error {
	// TODO: check rate-limit error for shopee
	if err.Status == http.StatusTooManyRequests {
		f, _ := strconv.ParseFloat(r.Header.Get("Retry-After"), 64)
		return RateLimitError{
			ResponseError: err,
			RetryAfter:    int(f),
		}
	}

	// if err.Status == http.StatusSeeOther {
	// todo
	// The response to the request can be found under a different URL in the
	// Location header and can be retrieved using a GET method on that resource.
	// }

	if err.Status == http.StatusNotAcceptable {
		err.Message = http.StatusText(err.Status)
	}

	return err
}

// shopee error maybe return status=200
// eg. {"error":"error_incalid_category.","message":"Invalid category ID","request_id":"2069449bd255af166cb52b0e15189d6d"}
// {"error":"error_category_is_block.","message":"Category is restricted","request_id":"97994a47af37a22da79cb910bfd9841a"}
func CheckResponseError(r *http.Response) error {
	shopeeError := struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer func() {
		// already read out, reload for next process
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}()

	if len(bodyBytes) > 0 {
		err := json.Unmarshal(bodyBytes, &shopeeError)
		if err != nil {
			return ResponseDecodingError{
				Body:    bodyBytes,
				Message: err.Error(),
				Status:  r.StatusCode,
			}
		}
	}

	if shopeeError.Error == "" && http.StatusOK <= r.StatusCode && r.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	responseError := ResponseError{
		Status:  r.StatusCode,
		Message: fmt.Sprintf("shopee-%s [%s]", shopeeError.Error, shopeeError.Message),
	}

	return wrapSpecificError(r, responseError)
}

// CreateAndDo performs a web request to Shopify with the given method (GET,
// POST, PUT, DELETE) and relative path (e.g. "/admin/orders.json").
// The data, options and resource arguments are optional and only relevant in
// certain situations.
// If the data argument is non-nil, it will be used as the body of the request
// for POST and PUT requests.
// The options argument is used for specifying request options such as search
// parameters like created_at_min
// Any data returned from Shopify will be marshalled into resource argument.
func (c *Client) CreateAndDo(method, relPath string, data, options, headers, resource interface{}) error {
	defer func() {
		// clear for next call
		c.ShopID = 0
		c.MerchantID = 0
		c.AccessToken = ""
	}()

	_, err := c.createAndDoGetHeaders(method, relPath, data, options, headers, resource)
	if err != nil {
		return err
	}
	return nil
}

// createAndDoGetHeaders creates an executes a request while returning the response headers.
func (c *Client) createAndDoGetHeaders(method, relPath string, data, options, headers, resource interface{}) (http.Header, error) {
	if strings.HasPrefix(relPath, "/") {
		// make sure it's a relative path
		relPath = strings.TrimLeft(relPath, "/")
	}

	relPath = path.Join("api/v2", relPath)

	if data != nil {
		params := data.(map[string]interface{})
		params["partner_id"] = c.app.PartnerID
		data = params
	}

	req, err := c.NewRequest(method, relPath, data, options, headers)
	if err != nil {
		return nil, err
	}

	return c.doGetHeaders(req, resource, false)
}

// Get performs a GET request for the given path and saves the result in the
// given resource.
func (c *Client) Get(path string, resource, options interface{}) error {
	return c.CreateAndDo("GET", path, nil, options, nil, resource)
}

// Post performs a POST request for the given path and saves the result in the
// given resource.
func (c *Client) Post(path string, data, resource interface{}) error {
	return c.CreateAndDo("POST", path, data, nil, nil, resource)
}

// Put performs a PUT request for the given path and saves the result in the
// given resource.
func (c *Client) Put(path string, data, resource interface{}) error {
	return c.CreateAndDo("PUT", path, data, nil, nil, resource)
}

// Delete performs a DELETE request for the given path
func (c *Client) Delete(path string) error {
	return c.CreateAndDo("DELETE", path, nil, nil, nil, nil)
}

// Upload performs a Upload request for the given path and saves the result in the
// given resource.
func (c *Client) Upload(relPath, fieldname, filename string, resource interface{}) error {
	req, err := c.NewfileUploadRequest(relPath, fieldname, filename)
	if err != nil {
		return err
	}

	if _, err := c.doGetHeaders(req, resource, true); err != nil {
		return err
	}
	return nil
}

// Creates a new file upload http request with optional extra params
func (c *Client) NewfileUploadRequest(relPath, paramName, filename string) (*http.Request, error) {
	if strings.HasPrefix(relPath, "/") {
		// make sure it's a relative path
		relPath = strings.TrimLeft(relPath, "/")
	}

	relPath = path.Join("api/v2", relPath)

	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)
	uri := u.String()

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, filepath.Base(filename))
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(part, file); err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)

	c.makeSignature(req)

	return req, nil
}
