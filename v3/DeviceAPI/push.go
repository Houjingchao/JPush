package DeviceAPI

import (
	"github.com/cocotyty/httpclient"
	"github.com/Houjingchao/JPush/v3/DeviceAPI/consts"
	"github.com/Houjingchao/JPush/v3/DeviceAPI/model"
)

//create by houjingchao on 17/08/10

type (
	Push interface {
		GetCids(count string) (model.CidResponse, error)
		PushAll() error
		PushByRegid(registrationID string) error
		PushByTag(tag string, context string) error
	}
	push struct {
		Authorization string
	}
)

func NewPush(authorization string) Push {
	return push{Authorization: authorization}
}

func (p push) PostClient(url string) *httpclient.HttpRequest {
	return httpclient.Post(url).
		Head("Authorization", p.Authorization).Head("Content-Type", "application/json")
}

func (p push) GetClient(url string, param string) *httpclient.HttpRequest {
	return httpclient.Get(url).
		Query(param, param).
		Head("Authorization", p.Authorization)
}

//默认count=1
func (p push) GetCids(count string) (model.CidResponse, error) {
	cidlist := model.CidResponse{}
	return cidlist, ResultGet(p.GetClient(consts.PushCidURL, count).Send(), &cidlist)
}

func (p push) PushAll() error {
	return nil
}

func (p push) PushByRegid(registrationID, title, context string, ) error {
	pushRequest := model.PushRequest{
		Platform: "all",
		Audience: model.Audience{
			RegistrationId: []string{registrationID},
		},
		Notification: model.Notification{
			Android: struct {
				Alert     string
				Title     string
				BuilderID int
				Extras    struct{ Newsid int`json:"newsid"` }
			}{Alert: title, Title: title },
			Ios: struct {
				Alert  string
				Sound  string
				Badge  string
				Extras struct{ Newsid int`json:"newsid"` }
			}{Alert: title },
		},
	}
	return ResultSet(p.PostClient(consts.PushURL).JSON(pushRequest).Send())
}

func (p push) PushByTag(tag, title, context string) error {
	pushRequest := model.PushRequest{
		Platform: "all",
		Audience: model.Audience{
			Tag: []string{tag},
		},
		Notification: model.Notification{
			Android: struct {
				Alert     string
				Title     string
				BuilderID int
				Extras    struct{ Newsid int`json:"newsid"` }
			}{Alert: title, Title: title},
			Ios: struct {
				Alert  string
				Sound  string
				Badge  string
				Extras struct{ Newsid int`json:"newsid"` }
			}{Alert: title },
		},
	}
	return ResultSet(p.PostClient(consts.PushURL).JSON(pushRequest).Send())
}
