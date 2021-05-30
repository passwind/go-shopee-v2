package goshopee

type ShopService interface {
	GetShopInfo (uint64, string) (*GetShopInfoResponse, error)
	GetProfile (uint64, string) (*GetProfileResponse, error)
}

type ShopServiceOp struct {
	client *Client
}

type ShopInfo struct {
	ShopName string `json:"shop_name"`
	Region string `json:"region"`
	Status string `json:"status"`
	SIPAffiShops []SIPAffiShops `json:"sip_affi_shops"`
	IsCB bool `json:"is_cb"`
	IsCNSC bool `json:"is_cnsc"`
}

type SIPAffiShops struct {
	AffiShopID uint64 `json:"affi_shop_id"`
	Region string `json:"region"`
}

type GetShopInfoResponse struct {
	ShopInfo

	RequestID string `json:"request_id"`
	AuthTime int64 `json:"auth_time"`
	ExpireTime int64 `json:"expire_time"`
}

func (s *ShopServiceOp)GetShopInfo (sid uint64, tok string) (*GetShopInfoResponse, error){
	path := "shop/get_shop_info"

	resp := new(GetShopInfoResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, nil)
	return resp, err
}

type GetProfileResponse struct {
	BaseResponse

	Response ShopProfile `json:"response"`
}

type ShopProfile struct {
	ShopLogo string `json:"shop_logo"`
	Description string `json:"description"`
	ShopName string `json:"shop_name"`
}

func (s *ShopServiceOp)GetProfile (sid uint64, tok string) (*GetProfileResponse, error){
	path := "shop/get_profile"

	resp := new(GetProfileResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, nil)
	return resp, err
}