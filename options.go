package goshopee

import (
	"net/http"
	"net/url"
)

// Option is used to configure client with options
type Option func(c *Client)

func WithRetry(retries int) Option {
	return func(c *Client) {
		c.retries = retries
	}
}

func WithLogger(logger LeveledLoggerInterface) Option {
	return func(c *Client) {
		c.log = logger
	}
}

func WithProxy(proxyHost string) Option {
	return func(c *Client) {
		proxyURL, err := url.Parse(proxyHost)
		if err != nil {
			return
		}
		c.Client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
}
