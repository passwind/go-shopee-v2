package goshopee

type DiscountService interface{
	DeleteDiscountItem(uint64, uint64, uint64, uint64, string) (*DeleteDiscountItemResponse, error)
}

type DiscountServiceOp struct {
	client *Client
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