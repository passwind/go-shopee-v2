package goshopee

import "testing"

func Test_GetAuthURL(t *testing.T) {
	setup()
	defer teardown()

	authURL,_:=client.Auth.GetAuthURL()
	t.Logf("auth url: %s",authURL)
} 