package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/options"
	"main/util"
	"testing"
	"time"
)

func TestOpenID(t *testing.T) {
	api := fmt.Sprintf("%s/getOpenId", domain)
	now := time.Now().UnixNano() / 1e6
	nonce := util.RandStringBytes(16)
	kgID := "test"
	appID := "4efab75f945745058c63d1b455df22b7"
	reqString := fmt.Sprintf("kgId=%s&appId=%s", kgID, appID)
	o := options.Option{
		AppID:     "4efab75f945745058c63d1b455df22b7",
		SecretKey: "xqHPeDQqlkOsbFIM",
	}
	checkSum := o.GetCheckSum(now, nonce, reqString, "")

	params := make(map[string]string, 2)
	params["kgId"] = kgID
	params["appId"] = appID

	headers := make(map[string]string, 4)
	headers["SAppId"] = o.AppID
	headers["time"] = string(now)
	headers["nonce"] = nonce
	headers["checkSum"] = checkSum

	resp, err := util.Get(api, params, headers)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%s", body)
	t.Log(string(body))
}
