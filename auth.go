package goshopee

import "fmt"

type AuthService interface {
	GetAuthURL() (string,error)
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
