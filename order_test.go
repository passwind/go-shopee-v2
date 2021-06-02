package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_GetOrderDetail(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/order/get_order_detail",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_order_detail_resp.json")))

	res,err:=client.Order.GetOrderDetail(shopID,[]string{"SN123"},nil,accessToken)
	if err!=nil {
		t.Errorf("Order.GetOrderDetail error: %s",err)
	}

	t.Logf("Order.GetOrderDetail: %#v",res)

	var expected string = "61630084074470"
	if res.Response.OrderList[0].PackageList[0].PackageNumber != expected {
		t.Errorf("PackageNumber returned %+v, expected %+v",res.Response.OrderList[0].PackageList[0].PackageNumber , expected)
	}
}