package goshopee

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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
	fmt.Println("debug: ",baseStr)
	h := hmac.New(sha256.New, []byte(s.client.app.PartnerKey))
	h.Write([]byte(baseStr))
	result := hex.EncodeToString(h.Sum(nil))
	return result,ts,nil
}