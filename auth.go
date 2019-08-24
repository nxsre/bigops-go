package bigops

import (
	"errors"
	ireq "github.com/soopsio/req"
)

func (c *Client) Authentication(user, pass string) (error) {
	resp, err := c.Req.Post(BASEURL+"/bastion/webapi/login", ireq.Param{
		"account": user,
		"pass":    pass,
	})

	if err != nil {
		return err
	}

	var data = BaseData{}
	resp.ToJSON(&data)

	if data.Code != 0 {
		return errors.New(data.Message)
	}
	return nil
}
