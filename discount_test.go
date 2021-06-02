package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_DeleteDiscountItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/delete_discount_item",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("delete_discount_item_resp.json")))

	res,err:=client.Discount.DeleteDiscountItem(shopID,1000029882,1776783,1467683,accessToken)
	if err!=nil {
		t.Errorf("Discount.DeleteDiscountItem error: %s",err)
	}

	t.Logf("Discount.DeleteDiscountItem: %#v",res)

	var expected string = "time error"
	if res.Response.ErrorList[0].FailMessage != expected {
		t.Errorf("FailMessage returned %+v, expected %+v",res.Response.ErrorList[0].FailMessage , expected)
	}
}