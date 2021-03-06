package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_GetCategory(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_category", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("category_list.json")))

	res, err := client.Product.GetCategory(shopID, "zh-hant", accessToken)
	if err != nil {
		t.Errorf("Product.GetCategory error: %s", err)
	}

	t.Logf("Product.GetCategory: %#v", res)

	var expectedID uint64 = 123
	if res.Response.CategoryList[0].CategoryID != expectedID {
		t.Errorf("CategoryID returned %+v, expected %+v", res.Response.CategoryList[0].CategoryID, expectedID)
	}
}

func Test_GetBrandList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_brand_list", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("brand_list.json")))

	res, err := client.Product.GetBrandList(shopID, 123, 0, 10, 1, accessToken)
	if err != nil {
		t.Errorf("Product.GetBrandList error: %s", err)
	}

	t.Logf("Product.GetBrandList: %#v", res)

	var expectedID uint64 = 2500139861
	if res.Response.BrandList[0].BrandID != expectedID {
		t.Errorf("BrandID returned %+v, expected %+v", res.Response.BrandList[0].BrandID, expectedID)
	}
}

func Test_GetDTSLimit(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_dts_limit", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("dts_limit.json")))

	res, err := client.Product.GetDTSLimit(shopID, 123, accessToken)
	if err != nil {
		t.Errorf("Product.GetDTSLimit error: %s", err)
	}

	t.Logf("Product.GetDTSLimit: %#v", res)

	var expected int = 7
	if res.Response.DaysToShipLimit.MaxLimit != expected {
		t.Errorf("DaysToShipLimit.MaxLimit returned %+v, expected %+v", res.Response.DaysToShipLimit.MaxLimit, expected)
	}
}

func Test_GetAttributes(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_attributes", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("attributes.json")))

	res, err := client.Product.GetAttributes(shopID, 123, "en", accessToken)
	if err != nil {
		t.Errorf("Product.GetAttributes error: %s", err)
	}

	t.Logf("Product.GetAttributes: %#v", res)

	var expectedID uint64 = 123
	if res.Response.AttributeList[0].AttributeID != expectedID {
		t.Errorf("AttributeList[0].AttributeID returned %+v, expected %+v", res.Response.AttributeList[0].AttributeID, expectedID)
	}

	var expectedBrandID uint64 = 2134
	if res.Response.AttributeList[0].AttributeValueList[0].ParentBrandList[0].ParentBrandID != expectedBrandID {
		t.Errorf("AttributeList[0].AttributeValueList[0].ParentBrandList[0].ParentBrandID returned %+v, expected %+v", res.Response.AttributeList[0].AttributeValueList[0].ParentBrandList[0].ParentBrandID, expectedBrandID)
	}
}

func Test_SupportSizeChart(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/support_size_chart", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("support_size_chart.json")))

	res, err := client.Product.SupportSizeChart(shopID, 123, accessToken)
	if err != nil {
		t.Errorf("Product.SupportSizeChart error: %s", err)
	}

	t.Logf("Product.SupportSizeChart: %#v", res)

	var expected bool = false
	if res.Response.SupportSizeChart != expected {
		t.Errorf("SupportSizeChart returned %+v, expected %+v", res.Response.SupportSizeChart, expected)
	}
}

func Test_UpdateSizeChart(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_size_chart", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	res, err := client.Product.UpdateSizeChart(shopID, 123, "test1234", accessToken)
	if err != nil {
		t.Errorf("Product.UpdateSizeChart error: %s", err)
	}

	t.Logf("Product.UpdateSizeChart: %#v", res)

	var expected string = "f634ea27eff8461b8f6f9ffa1d7ddab2"
	if res.RequestID != expected {
		t.Errorf("RequestID returned %+v, expected %+v", res.RequestID, expected)
	}
}

func Test_AddItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/add_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("add_item_resp.json")))

	var req AddItemRequest
	loadMockData("add_item_req.json", &req)

	res, err := client.Product.AddItem(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.AddItem error: %s", err)
	}

	t.Logf("Product.AddItem: %#v", res)

	var expectedID uint64 = 3000142341
	if res.Response.ItemID != expectedID {
		t.Errorf("ItemID returned %+v, expected %+v", res.Response.ItemID, expectedID)
	}
}

func Test_InitTierVariation(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/init_tier_variation", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("init_tier_variation_resp.json")))

	var req InitTierVariationRequest
	loadMockData("init_tier_variation_req.json", &req)

	res, err := client.Product.InitTierVariation(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.InitTierVariation error: %s", err)
	}

	t.Logf("Product.InitTierVariation: %#v", res)

	var expectedID uint64 = 12345
	if res.Response.Model[0].ModelID != expectedID {
		t.Errorf("ModelID returned %+v, expected %+v", res.Response.Model[0].ModelID, expectedID)
	}
}

func Test_AddModel(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/add_model", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("add_model_resp.json")))

	var req AddModelRequest
	loadMockData("add_model_req.json", &req)

	res, err := client.Product.AddModel(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.AddModel error: %s", err)
	}

	t.Logf("Product.AddModel: %#v", res)

	var expected float64 = 11.11
	if res.Response.Model[0].PriceInfo[0].OriginalPrice != expected {
		t.Errorf("OriginalPrice returned %+v, expected %+v", res.Response.Model[0].PriceInfo[0].OriginalPrice, expected)
	}
}

func Test_GetModelListt(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_model_list", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_model_list_resp.json")))

	res, err := client.Product.GetModelList(shopID, 123, accessToken)
	if err != nil {
		t.Errorf("Product.GetModelList error: %s", err)
	}

	t.Logf("Product.GetModelList: %#v", res)

	var expected uint64 = 2000458802
	if res.Response.Model[0].ModelID != expected {
		t.Errorf("ModelID returned %+v, expected %+v", res.Response.Model[0].ModelID, expected)
	}
}

func Test_GetItemBaseInfo(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_base_info", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_item_base_info_resp.json")))

	res, err := client.Product.GetItemBaseInfo(shopID, []uint64{123, 356}, accessToken)
	if err != nil {
		t.Errorf("Product.GetItemBaseInfo error: %s", err)
	}

	t.Logf("Product.GetItemBaseInfo: %#v", res)

	var expected string = "1e076dff0699d8e778c06dd6c02df1fe"

	if res.Response.ItemList[0].Image.ImageIDList[0] != expected {
		t.Errorf("Image ID returned %+v, expected %+v", res.Response.ItemList[0].Image.ImageIDList[0], expected)
	}
}

func Test_DeleteItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/delete_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	res, err := client.Product.DeleteItem(shopID, 34001, accessToken)
	if err != nil {
		t.Errorf("Product.DeleteItem error: %s", err)
	}

	t.Logf("Product.DeleteItem: %#v", res)

	var expected string = "f634ea27eff8461b8f6f9ffa1d7ddab2"
	if res.RequestID != expected {
		t.Errorf("res.RequestID returned %+v, expected %+v", res.RequestID, expected)
	}
}

func Test_UpdateItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("update_item_resp.json")))

	var req UpdateItemRequest
	loadMockData("update_item_req.json", &req)

	res, err := client.Product.UpdateItem(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UpdateItem error: %s", err)
	}

	t.Logf("Product.UpdateItem: %#v", res)

	var expected string = "Singpost - Registered Mail"
	if res.Response.LogisticInfo[1].LogisticName != expected {
		t.Errorf("LogisticName returned %+v, expected %+v", res.Response.LogisticInfo[1].LogisticName, expected)
	}
}

func Test_UnlistItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/unlist_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("unlist_item_resp.json")))

	var req UnlistItemRequest
	loadMockData("unlist_item_req.json", &req)

	res, err := client.Product.UnlistItem(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UnlistItem error: %s", err)
	}

	t.Logf("Product.UnlistItem: %#v", res)

	var expected string = "Can't unlist item when item is under promotion"
	if res.Response.FailureList[0].FailedReason != expected {
		t.Errorf("FailureList[0].FailedReason returned %+v, expected %+v", res.Response.FailureList[0].FailedReason, expected)
	}
}

func Test_DeleteModel(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/delete_model", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("delete_model_resp.json")))

	res, err := client.Product.DeleteModel(shopID, 34001, 123, accessToken)
	if err != nil {
		t.Errorf("Product.DeleteModel error: %s", err)
	}

	t.Logf("Product.DeleteModel: %#v", res)

	var expected string = "aaaaaaa"
	if res.RequestID != expected {
		t.Errorf("res.RequestID returned %+v, expected %+v", res.RequestID, expected)
	}
}

func Test_UpdateModel(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_model", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	var req UpdateModelRequest
	loadMockData("update_model_req.json", &req)

	res, err := client.Product.UpdateModel(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UpdateModel error: %s", err)
	}

	t.Logf("Product.UpdateModel: %#v", res)

	var expected string = "f634ea27eff8461b8f6f9ffa1d7ddab2"
	if res.RequestID != expected {
		t.Errorf("res.RequestID returned %+v, expected %+v", res.RequestID, expected)
	}
}

func Test_UpdatePrice(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_price", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("update_price_resp.json")))

	var req UpdatePriceRequest
	loadMockData("update_price_req.json", &req)

	res, err := client.Product.UpdatePrice(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UpdatePrice error: %s", err)
	}

	t.Logf("Product.UpdatePrice: %#v", res)

	var expected float64 = 11.11
	if res.Response.SuccessList[0].OriginalPrice != expected {
		t.Errorf("OriginalPrice returned %+v, expected %+v", res.Response.SuccessList[0].OriginalPrice, expected)
	}
}

func Test_UpdateStock(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_stock", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("update_stock_resp.json")))

	var req UpdateStockRequest
	loadMockData("update_stock_req.json", &req)

	res, err := client.Product.UpdateStock(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UpdateStock error: %s", err)
	}

	t.Logf("Product.UpdateStock: %#v", res)

	var expected int = 100
	if res.Response.SuccessList[0].NormalStock != expected {
		t.Errorf("NormalStock returned %+v, expected %+v", res.Response.SuccessList[0].NormalStock, expected)
	}
}

func Test_CategoryRecommend(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/category_recommend", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("category_recommend_resp.json")))

	res, err := client.Product.CategoryRecommend(shopID, "test", accessToken)
	if err != nil {
		t.Errorf("Product.CategoryRecommend error: %s", err)
	}

	t.Logf("Product.CategoryRecommend: %#v", res)

	var expectedID uint64 = 1000734
	if res.Response.CategoryID[0] != expectedID {
		t.Errorf("CategoryID returned %+v, expected %+v", res.Response.CategoryID[0], expectedID)
	}
}

func Test_GetItemPromotion(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_promotion", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_item_promotion_resp.json")))

	res, err := client.Product.GetItemPromotion(shopID, []uint64{123}, accessToken)
	if err != nil {
		t.Errorf("Product.GetItemPromotion error: %s", err)
	}

	t.Logf("Product.GetItemPromotion: %#v", res)

	var expected float64 = 12.12
	if res.Response.SuccessList[0].Promotion[0].PromotionPriceInfo[0].PromotionPrice != expected {
		t.Errorf("PromotionPrice returned %+v, expected %+v", res.Response.SuccessList[0].Promotion[0].PromotionPriceInfo[0].PromotionPrice, expected)
	}
}

func Test_UpdateTierVariation(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_tier_variation", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	var req UpdateTierVariationRequest
	loadMockData("update_tier_variation_req.json", &req)

	res, err := client.Product.UpdateTierVariation(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UpdateTierVariation error: %s", err)
	}

	t.Logf("Product.UpdateTierVariation: %#v", res)

	var expected string = "f634ea27eff8461b8f6f9ffa1d7ddab2"
	if res.RequestID != expected {
		t.Errorf("RequestID returned %+v, expected %+v", res.RequestID, expected)
	}
}
