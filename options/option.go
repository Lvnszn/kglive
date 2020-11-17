package options

import (
	"fmt"
	"log"
	"main/util"
)

// Option .
type Option struct {
	AppID     string `json:"app_id"`
	SecretKey string `json:"secret_key"`
}

// ServerOption .
type ServerOption struct {
	Version string `json:"version"`
	Port    string `json:"port"`
}

// GetCheckSum .
func (o Option) GetCheckSum(now, nonce, urlParams, queryBody string) (checkSum string) {
	checkSum = fmt.Sprintf("%s%s%s%s%s%s", o.AppID, now, nonce, urlParams, queryBody, o.SecretKey)
	log.Printf("AuthCode generated <%s>", checkSum)
	checkSum = util.MD5Encode(checkSum)
	return
}
