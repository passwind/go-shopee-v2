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

	res,err:=client.Product.GetCategory(shopID,accessToken,"zh-hant")
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