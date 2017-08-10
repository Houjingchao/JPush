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
		PushByRegid(registrationID, title, context string, extra map[string]string) error
		PushByTag(tag, title, context string, extra map[string]string) error
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

func (p push) PushByRegid(registrationID, title, context string, extra map[string]string) error {
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
				BigText   string
				Extras struct {
					Action  string`json:"action"`;
					Collect string`json:"collect"`;
					Func    string`json:"func"`;
					Url     string`json:"url"`
				}
			}{Alert: title, Title: title, BigText: context, Extras: struct {
				Action  string
				Collect string
				Func    string
				Url     string
			}{Action: extra["action"], Collect: extra["collect"], Func: extra["func"], Url: extra["url"]}},

			Ios: struct {
				Alert          string
				Sound          string
				Badge          string
				MutableContent string
				Extras struct {
					Action  string`json:"action"`;
					Collect string`json:"collect"`;
					Func    string`json:"func"`;
					Url     string`json:"url"`
				}
			}{Alert: title, MutableContent: context, Extras: struct {
				Action  string
				Collect string
				Func    string
				Url     string
			}{Action: extra["action"], Collect: extra["collect"], Func: extra["func"], Url: extra["url"]}},
		},
	}
	return ResultSet(p.PostClient(consts.PushURL).JSON(pushRequest).Send())
}

func (p push) PushByTag(tag, title, context string, extra map[string]string) error {
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
				BigText   string
				Extras struct {
					Action  string`json:"action"`;
					Collect string`json:"collect"`;
					Func    string`json:"func"`;
					Url     string`json:"url"`
				}
			}{Alert: title, Title: title, BigText: context, Extras: struct {
				Action  string
				Collect string
				Func    string
				Url     string
			}{Action: extra["action"], Collect: extra["collect"], Func: extra["func"], Url: extra["url"]}},

			Ios: struct {
				Alert          string
				Sound          string
				Badge          string
				MutableContent string
				Extras struct {
					Action  string`json:"action"`;
					Collect string`json:"collect"`;
					Func    string`json:"func"`;
					Url     string`json:"url"`
				}
			}{Alert: title, MutableContent: context, Extras: struct {
				Action  string
				Collect string
				Func    string
				Url     string
			}{Action: extra["action"], Collect: extra["collect"], Func: extra["func"], Url: extra["url"]}},
		},
	}
	return ResultSet(p.PostClient(consts.PushURL).JSON(pushRequest).Send())
}
