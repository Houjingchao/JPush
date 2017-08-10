package DeviceAPI

import (
	"github.com/cocotyty/httpclient"
	"bytes"
	"qiniupkg.com/x/errors.v7"
	"encoding/json"
)

/**
一个有返回值 一个没有返回值
 */
func ResultGet(resq *httpclient.HttpResponse, dest interface{}) error {
	body, err := resq.Body()
	if err != nil {
		return err
	}
	if len(body) != 0 && bytes.Contains(body, []byte("error")) {
		return errors.New(string(body))
	}
	return json.Unmarshal(body, dest)
}

func ResultSet(resp *httpclient.HttpResponse) error {
	body, err := resp.Body()
	if err != nil {
		return err
	} else if len(body) != 0 {
		return errors.New(string(body))
	}
	return nil
}
