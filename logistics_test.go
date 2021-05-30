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
