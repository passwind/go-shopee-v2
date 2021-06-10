package goshopee

//
type MediaSpaceService interface {
	UploadImage(string) (*UploadImageResponse,error)
}

// https://open.shopee.com/documents?module=91&type=1&id=660&version=2
type UploadImageResponse struct {
	BaseResponse

	Response UploadImageResponseData `json:"response"`
}

type UploadImageResponseData struct {
	ImageInfo ImageInfo `json:"image_info"`
}

type ImageInfo struct {
	ImageID string `json:"image_id"`
	ImageURLList []ImageURL `json:"image_url_list"`
}

type ImageURL struct {
	ImageURLRegion string `json:"image_url_region"`
	ImageURL string `json:"image_url"`
}

type MediaSpaceServiceOp struct {
	client *Client
}

func (s *MediaSpaceServiceOp)UploadImage(filename string) (*UploadImageResponse,error){
	path := "/media_space/upload_image"
	
	resp := new(UploadImageResponse)
	err := s.client.Upload(path, "image", filename, resp)
	return resp, err
}