package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_GetShopListByMerchant(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/merchant/get_shop_list_by_merchant", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_shop_list_by_merchant_resp.json")))

	res, err := client.Merchant.GetShopListByMerchant(merchantID, 1, 100, accessToken)
	if err != nil {
		t.Errorf("Merchant.GetShopListByMerchant error: %s", err)
	}

	t.Logf("Merchant.GetShopListByMerchant: %#v", res)

	var expectedID uint64 = 601306294
	if res.ShopList[0].ShopID != expectedID {
		t.Errorf("ShopID returned %+v, expected %+v", res.ShopList[0].ShopID, expectedID)
	}
}
