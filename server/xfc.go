package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/core"
	"main/options"
	"main/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// kugou doc: https://fanxing.kugou.com/open/doc/pages/develop/openapi/example.html
const (
	domain = "https://fx.service.kugou.com/fxservice/miniprogram/open"
)

// Server .
type Server struct {
	engine *gin.Engine
	config options.ServerOption
	o      options.Option
}

// Start run server
func (s *Server) Start() {
	v1 := s.engine.Group("/api/v1")
	v1.GET("/version", s.version)
	v1.GET("/getOpenId", s.getOpenID)
	v1.POST("/sendSysChatMsg", s.sendSysChatMsg)

	s.engine.Run(":8080")
}

func (s *Server) version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": s.config.Version,
	})
}

func (s *Server) getOpenID(c *gin.Context) {
	api := fmt.Sprintf("%s/getOpenId", domain)
	now := string(time.Now().UnixNano() / 1e6)
	nonce := util.RandStringBytes(16)
	kgID := c.Query("kgId")
	appID := c.Query("appId")
	reqString := fmt.Sprintf("kgId=%s&appId=%s", kgID, appID)
	checkSum := s.o.GetCheckSum(now, nonce, reqString, "")

	params := make(map[string]string, 2)
	params["kgId"] = kgID
	params["appId"] = appID

	headers := make(map[string]string, 4)
	headers["SAppId"] = s.o.AppID
	headers["time"] = now
	headers["nonce"] = nonce
	headers["checkSum"] = checkSum

	resp, err := util.Get(api, params, headers)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": err,
		})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"openID": string(body),
	})
}

func (s *Server) sendSysChatMsg(c *gin.Context) {
	api := fmt.Sprintf("%s/sendSysChatMsg", domain)
	now := time.Now().UnixNano() / 1e6
	nonce := util.RandStringBytes(16)
	starID := c.Query("starId")
	content := c.Query("content")

	cr := &core.ContentReq{
		Content: content,
		SessionContext: core.SessionContext{
			ID:     util.RandStringBytes(30),
			Time:   int(now),
			AppID:  s.o.AppID,
			StarID: starID,
		},
	}

	dd, err := json.Marshal(cr)
	if err != nil {
		return
	}

	checkSum := s.o.GetCheckSum(string(now), nonce, "", string(dd))

	params := make(map[string]string, 5)
	params["id"] = cr.ID
	params["content"] = content
	params["starId"] = starID
	params["time"] = string(now)
	params["appId"] = s.o.AppID

	headers := make(map[string]string, 4)
	headers["SAppId"] = s.o.AppID
	headers["time"] = string(now)
	headers["nonce"] = nonce
	headers["checkSum"] = checkSum

	resp, err := util.Post(api, nil, params, headers)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": err,
		})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"openID": string(body),
	})
}

// New .
func New(o options.Option) *Server {
	s := new(Server)
	s.engine = gin.Default()
	s.o = o

	return s
}
