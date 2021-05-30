package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_GetShopInfo(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/shop/get_shop_info",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_shop_info_resp.json")))

	res,err:=client.Shop.GetShopInfo(shopID,accessToken)
	if err!=nil {
		t.Errorf("Shop.GetShopInfo error: %s",err)
	}

	t.Logf("Shop.GetShopInfo: %#v",res)

	var expectedID uint64 = 261373
	if res.SIPAffiShops[0].AffiShopID != expectedID {
		t.Errorf("SIPAffiShops[0].AffiShopID returned %+v, expected %+v",res.SIPAffiShops[0].AffiShopID , expectedID)
	}
}