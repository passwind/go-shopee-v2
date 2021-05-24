package goshopee

import "testing"

func Test_GetAuthURL(t *testing.T) {
	setup()
	defer teardown()

	authURL,_:=client.Auth.GetAuthURL()
	t.Logf("auth url: %s",authURL)
} 

func Test_GetCancelAuthURL(t *testing.T) {
	setup()
	defer teardown()

	authURL,_:=client.Auth.GetCancelAuthURL()
	t.Logf("cancel auth url: %s",authURL)
} 