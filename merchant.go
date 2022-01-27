package goshopee

// https://open.shopee.com/documents/v2/v2.merchant.get_shop_list_by_merchant?module=93&type=1
type MerchantService interface {
	// GetMerchantInfo() (*GetMerchantInfoResponse, error)
	GetShopListByMerchant(mid uint64, pageNo, pageSize int, tok string) (*GetShopListByMerchantResponse, error)
}

type GetShopListByMerchantResponse struct {
	BaseResponse

	IsCNSC   bool                                `json:"is_cnsc"`
	ShopList []GetShopListByMerchantResponseData `json:"shop_list"`
	More     bool                                `json:"more"`
}

type GetShopListByMerchantResponseData struct {
	ShopID       uint64                                  `json:"shop_id"`
	SIPAffiShops []GetShopListByMerchantResponseDataShop `json:"sip_affi_shops,omitempty"`
	ShopIsCNSC   bool                                    `json:"shop_is_cnsc"`
}

type GetShopListByMerchantResponseDataShop struct {
	AffiShopID uint64 `json:"affi_shop_id"`
}

type MerchantServiceOp struct {
	client *Client
}

type GetShopListByMerchantRequest struct {
	PageNo   int `url:"page_no"`
	PageSize int `url:"page_size"`
	// MerchantID uint64 `url:"merchant_id"` add with token
}

func (s *MerchantServiceOp) GetShopListByMerchant(mid uint64, pageNo, pageSize int, tok string) (*GetShopListByMerchantResponse, error) {
	path := "/merchant/get_shop_list_by_merchant"

	opt := GetShopListByMerchantRequest{
		PageNo:   pageNo,
		PageSize: pageSize,
	}

	resp := new(GetShopListByMerchantResponse)
	err := s.client.WithMerchant(mid, tok).Get(path, resp, opt)
	return resp, err
}
