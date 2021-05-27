package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_GetCategory(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_category",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("category_list.json")))

	res,err:=client.Product.GetCategory(shopID,"zh-hant",accessToken)
	if err!=nil {
		t.Errorf("Product.GetCategory error: %s",err)
	}

	t.Logf("Product.GetCategory: %#v",res)

	var expectedID int64 = 123
	if res.Response.CategoryList[0].CategoryID != expectedID {
		t.Errorf("CategoryID returned %+v, expected %+v",res.Response.CategoryList[0].CategoryID , expectedID)
	}
}

func Test_GetBrandList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_brand_list",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("brand_list.json")))

	res,err:=client.Product.GetBrandList(shopID,123,0,10,1,accessToken)
	if err!=nil {
		t.Errorf("Product.GetBrandList error: %s",err)
	}

	t.Logf("Product.GetBrandList: %#v",res)

	var expectedID int64 = 2500139861
	if res.Response.BrandList[0].BrandID != expectedID {
		t.Errorf("BrandID returned %+v, expected %+v",res.Response.BrandList[0].BrandID , expectedID)
	}
}

func Test_GetDTSLimit(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_dts_limit",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("dts_limit.json")))

	res,err:=client.Product.GetDTSLimit(shopID,123,accessToken)
	if err!=nil {
		t.Errorf("Product.GetDTSLimit error: %s",err)
	}

	t.Logf("Product.GetDTSLimit: %#v",res)

	var expected int = 7
	if res.Response.DaysToShipLimit.MaxLimit != expected {
		t.Errorf("DaysToShipLimit.MaxLimit returned %+v, expected %+v",res.Response.DaysToShipLimit.MaxLimit , expected)
	}
}

func Test_GetAttributes(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_attributes",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("attributes.json")))

	res,err:=client.Product.GetAttributes(shopID,123,"en",accessToken)
	if err!=nil {
		t.Errorf("Product.GetAttributes error: %s",err)
	}

	t.Logf("Product.GetAttributes: %#v",res)

	var expectedID int64 = 123
	if res.Response.AttributeList[0].AttributeID != expectedID {
		t.Errorf("AttributeList[0].AttributeID returned %+v, expected %+v",res.Response.AttributeList[0].AttributeID , expectedID)
	}

	var expectedBrandID int64 = 2134
	if res.Response.AttributeList[0].AttributeValueList[0].ParentBrandList[0].ParentBrandID != expectedBrandID {
		t.Errorf("AttributeList[0].AttributeValueList[0].ParentBrandList[0].ParentBrandID returned %+v, expected %+v",res.Response.AttributeList[0].AttributeValueList[0].ParentBrandList[0].ParentBrandID , expectedBrandID)
	}
}

func Test_SupportSizeChart(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/support_size_chart",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("support_size_chart.json")))

	res,err:=client.Product.SupportSizeChart(shopID,123,accessToken)
	if err!=nil {
		t.Errorf("Product.SupportSizeChart error: %s",err)
	}

	t.Logf("Product.SupportSizeChart: %#v",res)

	var expected bool = false
	if res.Response.SupportSizeChart != expected {
		t.Errorf("SupportSizeChart returned %+v, expected %+v",res.Response.SupportSizeChart , expected)
	}
}

func Test_UpdateSizeChart(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_size_chart",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	res,err:=client.Product.UpdateSizeChart(shopID,123,"test1234",accessToken)
	if err!=nil {
		t.Errorf("Product.UpdateSizeChart error: %s",err)
	}

	t.Logf("Product.UpdateSizeChart: %#v",res)

	var expected string = "f634ea27eff8461b8f6f9ffa1d7ddab2"
	if res.RequestID != expected {
		t.Errorf("RequestID returned %+v, expected %+v",res.RequestID , expected)
	}
}

func Test_AddItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/add_item",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("add_item_resp.json")))

	var req AddItemRequest
	loadMockData("add_item_req.json",&req)

	res,err:=client.Product.AddItem(shopID,req,accessToken)
	if err!=nil {
		t.Errorf("Product.AddItem error: %s",err)
	}

	t.Logf("Product.AddItem: %#v",res)

	var expectedID int64 = 3000142341
	if res.Response.ItemID != expectedID {
		t.Errorf("ItemID returned %+v, expected %+v", res.Response.ItemID , expectedID)
	}
}

func Test_InitTierVariation(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/init_tier_variation",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("init_tier_variation_resp.json")))

	var req InitTierVariationRequest
	loadMockData("init_tier_variation_req.json",&req)

	res,err:=client.Product.InitTierVariation(shopID,req,accessToken)
	if err!=nil {
		t.Errorf("Product.InitTierVariation error: %s",err)
	}

	t.Logf("Product.InitTierVariation: %#v",res)

	var expectedID int64 = 12345
	if res.Response.Model[0].ModelID != expectedID {
		t.Errorf("ModelID returned %+v, expected %+v", res.Response.Model[0].ModelID , expectedID)
	}
}

func Test_AddModel(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/add_model",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("add_model_resp.json")))

	var req AddModelRequest
	loadMockData("add_model_req.json",&req)

	res,err:=client.Product.AddModel(shopID,req,accessToken)
	if err!=nil {
		t.Errorf("Product.AddModel error: %s",err)
	}

	t.Logf("Product.AddModel: %#v",res)

	var expected float64 = 11.11
	if res.Response.Model[0].PriceInfo[0].OriginalPrice != expected {
		t.Errorf("OriginalPrice returned %+v, expected %+v", res.Response.Model[0].PriceInfo[0].OriginalPrice , expected)
	}
}