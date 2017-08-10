package DeviceAPI

import (
	"github.com/cocotyty/httpclient"
	"github.com/Houjingchao/JPush/v3/DeviceAPI/consts"
)

//create by houjingchao on 17/08/10

type (
	Push interface {
		GetCids(count, ty string) (string, error)
		PushAll()
		PushByID()
		PushByTag()
	}
	push struct {
		Authorization string
	}
)

func NewPush(authorization string) Push {
	return push{Authorization: authorization}
}

func (p push) PostClient(url, param string) *httpclient.HttpRequest {
	return httpclient.Post(url + param).Head("Authorization", p.Authorization)
}

func (p push) GetClient(url string, param string) *httpclient.HttpRequest {
	return httpclient.Get(url + param).Head("Authorization", p.Authorization)
}

func (p push) GetCids(count, ty string) (string, error) {

	return ResultGet(p.GetClient(consts.DevicesURL, reistrationID).Send(), &ds)
}

func (p push) PushAll() () {
}

func (p push) PushByID() () {
}

func (p push) PushByTag() () {
}
