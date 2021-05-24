package goshopee

import "fmt"

type AuthService interface {
	GetAuthURL() (string,error)
	GetCancelAuthURL() (string,error)
	GetToken(uint64,uint64,string) (*AccessTokenResponse,error)
}

type AccessTokenResponse struct {
	BaseResponse

	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn int `json:"expire_in"`
	MerchantIDList []uint64 `json:"merchant_id_list,omitempty"`
	ShopIDList []uint64 `json:"shop_id_list,omitempty"`
}

type AuthServiceOp struct {
	client *Client
}

func (s *AuthServiceOp)GetAuthURL() (string,error) {
	rurl := s.client.app.RedirectURL
	path:="/api/v2/shop/auth_partner"
	sign,ts,_ := s.client.Util.Sign(path)
	aurl := fmt.Sprintf("%s%s?partner_id=%d&timestamp=%d&sign=%s&redirect=%s", s.client.app.APIURL,path, s.client.app.PartnerID,ts,sign, rurl)
	return aurl,nil
}

func (s *AuthServiceOp)GetCancelAuthURL() (string,error) {
	rurl := s.client.app.RedirectURL
	path:="/api/v2/shop/cancel_auth_partner"
	sign,ts,_ := s.client.Util.Sign(path)
	aurl := fmt.Sprintf("%s%s?partner_id=%d&timestamp=%d&sign=%s&redirect=%s", s.client.app.APIURL,path, s.client.app.PartnerID,ts,sign, rurl)
	return aurl,nil
}

func (s *AuthServiceOp)GetToken(sid uint64, aid uint64, code string) (*AccessTokenResponse,error){
	path := "/auth/token/get"
	params := map[string]interface{}{
		"code": code,
	}
	if sid!=0{
		params["shop_id"]=sid
	}else if aid!=0{
		params["main_account_id"]=aid
	}

	resp := new(AccessTokenResponse)
	err := s.client.Post(path, params, resp)
	return resp, err
}
