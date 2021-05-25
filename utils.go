package goshopee

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type UtilService interface {
	Sign(string) (string, int64,error)
}

type UtilServiceOp struct {
	client *Client
}

func (s *UtilServiceOp)Sign(plainText string) (string,int64,error) {
	ts:=time.Now().Unix()
	baseStr := fmt.Sprintf("%d%s%d",s.client.app.PartnerID,plainText,ts)
	h := hmac.New(sha256.New, []byte(s.client.app.PartnerKey))
	h.Write([]byte(baseStr))
	result := hex.EncodeToString(h.Sum(nil))
	return result,ts,nil
}

func StructToMap(in interface{}) (map[string]interface{},error) {
	byts,err:=json.Marshal(in)
	if err!=nil {
		return nil,fmt.Errorf("error to perpare request body: %s", err)
	}
	var res map[string]interface{}
	if err:=json.Unmarshal(byts,&res);err!=nil {
		return nil,fmt.Errorf("error to perpare request body 1: %s", err)
	}
	return res,nil
}