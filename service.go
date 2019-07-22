package bigops

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"github.com/soopsio/bigops-go/utils"
	ireq "github.com/soopsio/req"
)

type Service struct {
	ServiceTreeId int    `json:"serviceTreeId"`
	ServiceName   string `json:"serviceName"`
	Num           int    `json:"num"`
}

func (c *Client) GetServices(accout, keyword string) ([]Service, error) {

	resp, err := c.Req.Get(BASEURL+"/bastion/webapi/servicelist", ireq.Param{
		"account": accout,
		//"keyword": keyword,
	})

	if err != nil {
		return nil, err
	}

	c.logger.Debug("service resp", zap.String("resp", resp.String()))

	var data = BaseData{}
	err = resp.ToJSON(&data)
	if err != nil {
		return nil, err
	}

	if data.Code != 0 {
		return nil, errors.New(data.Message)
	}

	key, _ := utils.AESSHA1PRNG([]byte(c.SecretKey), 128)
	encstr, err := base64.StdEncoding.DecodeString(data.Data)
	if err != nil {
		return nil, err
	}
	decstr := utils.Decrypt(encstr, key)
	c.logger.Debug(string(decstr))

	c.logger.Debug(string(utils.PKCS7UPad(decstr)))

	services := []Service{}
	err = json.Unmarshal(utils.PKCS7UPad(decstr), &services)
	if err != nil {
		return nil, err
	}

	return services, nil
}
