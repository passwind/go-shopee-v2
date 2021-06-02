package goshopee

type DiscountService interface{
	AddDiscount(uint64, AddDiscountRequest, string) (*AddDiscountResponse, error)
	AddDiscountItem(uint64, AddDiscountItemRequest, string) (*AddDiscountItemResponse, error)
	DeleteDiscountItem(uint64, uint64, uint64, uint64, string) (*DeleteDiscountItemResponse, error)
}

type DiscountServiceOp struct {
	client *Client
}

type AddDiscountRequest struct {
	DiscountName string `json:"discount_name"`
	StartTime int64 `json:"start_time"`
	EndTime int64 `json:"end_time"`
}

type AddDiscountResponse struct {
	BaseResponse

	Response AddDiscountResponseData `json:"response"`
}

type AddDiscountResponseData struct {
	DiscountID uint64 `json:"discount_id"`
}

func (s *DiscountServiceOp)AddDiscount(sid uint64, data AddDiscountRequest, tok string) (*AddDiscountResponse, error) {
	path := "/discount/add_discount"
	resp := new(AddDiscountResponse)
	req,err:=StructToMap(data)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

type AddDiscountItemRequest struct {
	DiscountID uint64 `json:"discount_id"`
	ItemList []AddDiscountItemRequestData `json:"item_list"`
}

type AddDiscountItemRequestData struct {
	ItemID uint64 `json:"item_id"`
	ModelList []AddDiscountItemRequestDataModel `json:"model_list"`
	ItemPromotionPrice float64 `json:"item_promotion_price"`
	PurchaseLimit int `json:"purchase_limit"`
	ItemPromotionStock int `json:"item_promotion_stock"`
}

type AddDiscountItemRequestDataModel struct {
	ModelID uint64 `json:"model_id"`
	ModelPromotionPrice float64 `json:"model_promotion_price"`
	ModelPromotionStock int `json:"model_promotion_stock"`
}

type AddDiscountItemResponse struct {
	BaseResponse

	Response AddDiscountItemResponseData `json:"response"`
}

type AddDiscountItemResponseData struct {
	DiscountID uint64 `json:"discount_id"`
	Count int `json:"count"`
	ErrorList []AddDiscountItemResponseDataError `json:"error_list"`
}

type AddDiscountItemResponseDataError struct {
	ItemID uint64 `json:"item_id"`
	ModelID uint64 `json:"model_id"`
	FailMessage string `json:"fail_message"`
	FailError string `json:"fail_error"`
}

func (s *DiscountServiceOp)AddDiscountItem(sid uint64, data AddDiscountItemRequest, tok string) (*AddDiscountItemResponse, error) {
	path := "/discount/add_discount_item"
	resp := new(AddDiscountItemResponse)
	req,err:=StructToMap(data)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

type DeleteDiscountItemResponse struct {
	BaseResponse

	Response DeleteDiscountItemResponseData `json:"response"`
}

type DeleteDiscountItemResponseData struct {
	DiscountID uint64 `json:"discount_id"`
	ErrorList []DeleteDiscountItemResponseDataError `json:"error_list"`
}

type DeleteDiscountItemResponseDataError struct {
	ItemID uint64 `json:"item_id"`
	ModelID uint64 `json:"model_id"`
	FailMessage string `json:"fail_message"`
	FailError string `json:"fail_error"`
}

func (s *DiscountServiceOp)DeleteDiscountItem(sid, discountID, itemID, modelID uint64, tok string) (*DeleteDiscountItemResponse, error) {
	path := "/discount/delete_discount_item"
	wrappedData := map[string]interface{}{
		"discount_id": discountID,
		"item_id": itemID,
		"model_id":  modelID,
	}
	resp := new(DeleteDiscountItemResponse)
	err := s.client.WithShop(sid,tok).Post(path, wrappedData, resp)
	return resp, err
}