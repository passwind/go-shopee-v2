package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_GetChannelList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_channel_list",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("channel_list.json")))

	res,err:=client.Logistics.GetChannelList(shopID,accessToken)
	if err!=nil {
		t.Errorf("Logistics.GetChannelList error: %s",err)
	}

	t.Logf("Logistics.GetChannelList: %#v",res)

	if len(res.Response.LogisticsChannelList)!=5 {
		t.Errorf("LogisticsChannelList len return %v, expected 5",len(res.Response.LogisticsChannelList))
	}
	var expectedID uint64 = 5116
	if res.Response.LogisticsChannelList[4].LogisticsChannelID != expectedID {
		t.Errorf("LogisticsChannelList[4].LogisticsChannelID returned %+v, expected %+v",res.Response.LogisticsChannelList[4].LogisticsChannelID , expectedID)
	}
}

func Test_GetShippingParameter(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_shipping_parameter",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_shipping_parameter_resp.json")))

	res,err:=client.Logistics.GetShippingParameter(shopID,"SN123",accessToken)
	if err!=nil {
		t.Errorf("Logistics.GetShippingParameter error: %s",err)
	}

	t.Logf("Logistics.GetShippingParameter: %#v",res)

	if len(res.Response.InfoNeeded.Pickup)!=2 {
		t.Errorf("LogisticsChannelList len return %v, expected 2",len(res.Response.InfoNeeded.Pickup))
	}
	var expected string = "hhh, #34"
	if res.Response.Pickup.AddressList[1].Address != expected {
		t.Errorf("Pickup.AddressList[1].Address returned %+v, expected %+v",res.Response.Pickup.AddressList[1].Address , expected)
	}
}

func Test_ShipOrder(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/ship_order",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	var req ShipOrderRequest
	loadMockData("ship_order_req.json",&req)
	
	res,err:=client.Logistics.ShipOrder(shopID,req,accessToken)
	if err!=nil {
		t.Errorf("Logistics.ShipOrder error: %s",err)
	}

	t.Logf("Logistics.ShipOrder: %#v",res)

	var expected string = "f634ea27eff8461b8f6f9ffa1d7ddab2"
	if res.RequestID != expected {
		t.Errorf("RequestID returned %+v, expected %+v",res.RequestID , expected)
	}
}
