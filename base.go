package bigops

import (
	"github.com/soopsio/bigops-go/req"
	"go.uber.org/zap"
)

const BASEURL = "http://work.bigops.com/api"

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

func NewClient(accessKey, secretKey string, debug bool, logger *zap.Logger) *Client {
	c := &Client{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Req:       req.New(accessKey, secretKey, debug, logger),
		logger:    logger,
	}
	return c
}
