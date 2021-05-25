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