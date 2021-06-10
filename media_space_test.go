package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_UploadImage(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media_space/upload_image",app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("upload_image.json")))

	res,err:=client.Media.UploadImage("fixtures/test.jpg")
	if err!=nil {
		t.Errorf("Media.UploadImage error: %s",err)
	}

	t.Logf("return image: %#v",res)

	var expectedID string = "e721546cbfafcb14ac6ae6c7cf57e455"
	if res.Response.ImageInfo.ImageID != expectedID {
		t.Errorf("ImageInfo.ImageID returned %+v, expected %+v", res.Response.ImageInfo.ImageID, expectedID)
	}
}