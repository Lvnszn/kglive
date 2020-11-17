package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/options"
	"main/util"
	"time"
)

var (
	id        = "4efab75f945745058c63d1b455df22b7"
	secretKey = "xqHPeDQqlkOsbFIM"
)

func main() {
	o := options.Option{
		AppID:     id,
		SecretKey: secretKey,
	}

	api := fmt.Sprintf("https://fx.service.kugou.com/fxservice/miniprogram/open/getOpenId")
	now := time.Now().UnixNano() / 1e6
	nonce := util.RandStringBytes(16)
	kgID := "1414978005"
	reqString := fmt.Sprintf("appId=%s&kgId=%s", kgID, o.AppID)

	checkSum := o.GetCheckSum(fmt.Sprint(now), nonce, reqString, "")
	params := make(map[string]string, 2)
	params["kgId"] = kgID
	params["appId"] = o.AppID

	headers := make(map[string]string, 4)
	headers["SAppId"] = o.AppID
	headers["time"] = fmt.Sprint(now)
	headers["nonce"] = nonce
	headers["checkSum"] = checkSum

	resp, err := util.Get(api, params, headers)
	if err != nil {
		print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	log.Printf("%s", body)

	// s := server.New(o)
	// s.Start()
}
