package bigops

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/soopsio/bigops-go/utils"
	ireq "github.com/soopsio/req"
)

type Session struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Num      int    `json:"num"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
	User     string `json:"user"`
	HostID   int    `json:"hostid"`
	Pass     string `json:"pass"`
	TreeName string `json:"treeName"`
}

func (c *Client) GetSession(accout string, serviceId int, protocols string, keyword string) ([]Session, error) {
	c.logger.Debug(accout + " " + fmt.Sprint(serviceId) + " " + protocols)

	resp, err := c.Req.Get(BASEURL+"/bastion/webapi/sessionlist", ireq.Param{
		"account":       accout,
		"serviceTreeId": serviceId,
		//"proto":         protocols,
		"keyword": keyword,
	})
	if err != nil {
		c.logger.Error(err.Error())
		return nil, err
	}
	c.logger.Debug(resp.String())

	var data = BaseData{}
	err = resp.ToJSON(&data)
	if err != nil {
		return nil, err
	}
	if data.Code != 0 {
		return nil, errors.New(data.Message)
	}

	key, _ := utils.AESSHA1PRNG([]byte(c.SecretKey), 128)
	encstr, _ := base64.StdEncoding.DecodeString(data.Data)
	decstr := utils.Decrypt(encstr, key)

	c.logger.Debug(string(utils.PKCS7UPad(decstr)))

	sessions := []Session{}
	err = json.Unmarshal(utils.PKCS7UPad(decstr), &sessions)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}
