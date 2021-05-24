package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

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

func Test_GetToken(t *testing.T) {
	setup()
	defer teardown()

	apiurl:=fmt.Sprintf("%s/api/v2/auth/token/get",app.APIURL)

	httpmock.RegisterResponder("POST", apiurl,
		httpmock.NewBytesResponder(200, loadFixture("access_token.json")))

	res,err:=client.Auth.GetToken(123456,0,"testcode")
	if err!=nil {
		t.Errorf("Auth.GetToken error: %s",err)
	}

	t.Logf("return tok: %#v",res)

	var expectedToken string = "accesstoken"
	if res.AccessToken != expectedToken {
		t.Errorf("Token.AccessToken returned %+v, expected %+v", res.AccessToken, expectedToken)
	}
}