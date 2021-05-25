package goshopee

type ProductService interface {
	GetCategory(int64, string, string) (*GetCategoryResponse,error)
	GetBrandList(int64, int64, int, int, int, string) (*GetBrandListResponse, error)
}

type GetCategoryResponse struct {
	BaseResponse

	Response GetCategoryResponseData `json:"response"`
}

type GetCategoryResponseData struct {
	CategoryList []Category `json:"category_list"`
}

type Category struct {
	CategoryID int64 `json:"category_id"`
	ParentCategoryID int64 `json:"parent_category_id"`
	OriginalCategoryName string `json:"original_category_name"`
	DisplayCategoryName string `json:"display_category_name"`
	HasChildren bool `json:"has_children"`
}

type ProductServiceOp struct {
	client *Client
}

type GetCategoryRequest struct {
	Language string `url:"language"`
}

type GetBrandListRequest struct {
	Offset int `url:"offset"`
	PageSize int `url:"page_size"`
	CategoryID int64 `url:"category_id"`
	Status int `url:"status"`
}

type GetBrandListResponse struct {
	BaseResponse

	Response GetBrandListResponseData `json:"response"`
}

type GetBrandListResponseData struct {
	BrandList []Brand `json:"brand_list"`
	HasNextPage bool `json:"has_next_page"`
	NextOffset int `json:"next_offset"`
	IsMandatory bool `json:"is_mandatory"`
	InputType string `json:"input_type"`
}

type Brand struct {
	BrandID int64 `json:"brand_id"`
	OriginalBrandName string `json:"original_brand_name"`
	DisplayBrandName string `json:"display_brand_name"`
}

func (s *ProductServiceOp)	GetCategory(sid int64, tok, lang string) (*GetCategoryResponse,error){
	path := "/product/get_category"

	opt:=GetCategoryRequest{lang}

	resp := new(GetCategoryResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp)	GetBrandList(sid, cid int64, status, offset, pageSize int, tok string) (*GetBrandListResponse, error){
	path := "/product/get_brand_list"

	opt:=GetBrandListRequest{
		CategoryID: cid,
		Offset: offset,
		PageSize: pageSize,
		Status: status,
	}

	resp := new(GetBrandListResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}