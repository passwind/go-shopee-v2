package goshopee

import "fmt"

// https://open.shopee.com/documents?module=87&type=2&id=58&version=2
type AuthService interface {
	GetAuthURL() (string,error)
	GetCancelAuthURL() (string,error)
	GetAccessToken(int64,int64,string) (*AccessTokenResponse,error)
	RefreshAccessToken(int64,int64,string) (*RefreshAccessTokenResponse,error)
}

type AccessTokenResponse struct {
	BaseResponse

	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn int `json:"expire_in"`
	MerchantIDList []int64 `json:"merchant_id_list,omitempty"`
	ShopIDList []int64 `json:"shop_id_list,omitempty"`
}

type RefreshAccessTokenResponse struct {
	BaseResponse

	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn int `json:"expire_in"`
	PartnerID int64 `json:"partner_id"`
	MerchantID int64 `json:"merchant_id"`
	ShopID int64 `json:"shop_id"`
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

func (s *AuthServiceOp)GetAccessToken(sid int64, aid int64, code string) (*AccessTokenResponse,error){
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

func (s *AuthServiceOp)RefreshAccessToken(sid int64, aid int64, refresh string) (*RefreshAccessTokenResponse,error){
	path := "/auth/access_token/get"
	params := map[string]interface{}{
		"refresh_token": refresh,
	}
	if sid!=0{
		params["shop_id"]=sid
	}else if aid!=0{
		params["main_account_id"]=aid
	}

	resp := new(RefreshAccessTokenResponse)
	err := s.client.Post(path, params, resp)
	return resp, err
}
