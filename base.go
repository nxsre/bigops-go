package bigops

import (
	"github.com/soopsio/bigops-go/req"
	"go.uber.org/zap"
	"fmt"
)

var BASEURL_P = "http://%s/api"
var BASEURL = ""

type BaseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type Client struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Req       *req.Req
	logger    *zap.Logger
}

func NewClient(domain, accessKey, secretKey string, debug bool, logger *zap.Logger) *Client {
	BASEURL = fmt.Sprintf(BASEURL_P, domain)
	c := &Client{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Req:       req.New(accessKey, secretKey, debug, logger),
		logger:    logger,
	}
	return c
}
