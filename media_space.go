package goshopee

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

//
type MediaSpaceService interface {
	UploadImage(string) (*UploadImageResponse,error)
}

// https://open.shopee.com/documents?module=91&type=1&id=660&version=2
type UploadImageResponse struct {
	BaseResponse

	Response *UploadImageResponseData `json:"response"`
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


	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open file error: %s [%s]",err,path)
	}
	defer file.Close()

	filebody := &bytes.Buffer{}
	writer := multipart.NewWriter(filebody)
	defer writer.Close()

	part, err := writer.CreateFormFile("image", filepath.Base(path))
	if err != nil {
		return nil, fmt.Errorf("prepare upload error: %s",err)
	}
	if _, err = io.Copy(part, file);err!=nil {
		return nil, fmt.Errorf("prepare upload 1 error: %s",err)
	}
	
	resp := new(UploadImageResponse)
	headers:=map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}
	err = s.client.Upload(path, filebody, headers, resp)
	return resp, err
}