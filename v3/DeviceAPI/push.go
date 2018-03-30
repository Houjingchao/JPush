package DeviceAPI

import (
	"github.com/cocotyty/httpclient"
	"github.com/Houjingchao/JPush/v3/DeviceAPI/consts"
	"github.com/Houjingchao/JPush/v3/DeviceAPI/model"
	"encoding/json"
	"errors"
	"strings"
	"strconv"
)

//create by houjingchao on 17/08/10

type Response struct {
	Sendno string `json:"sendno"`
	MsgId  string `json:"msg_id"`
}

type JpushError struct {
	Error struct {
		Message string `db:"message" json:"message"`
		Code int `db:"code" json:"code"`
	} `db:"error" json:"error"`
}

type (
	Push interface {
		GetCids(count string) (model.CidResponse, error)
		PushAll() error
		PushByRegid(registrationID, title, context string, extra model.Extras) error
		PushByTag(tag, title, context string, extra model.Extras) error
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
		Head("Authorization", p.Authorization)
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

func (p push) PushByRegid(registrationID, title, content string, extra model.Extras) error {
	pushRequest := model.PushRequest{
		Platform: "all",
		Audience: model.Audience{
			RegistrationId: []string{registrationID},
		},
		Notification: model.Notification{
			Android: model.Android{
				Alert:   content,
				Title:   title,
				BigText: "",
				Extras:  extra,
			},
			Ios: model.Ios{
				Alert:  title+"  "+content,
				Extras: extra,
			},
		},
	}
	response := &Response{}
	body, err := p.PostClient(consts.PushURL).JSON(pushRequest).Send().Body()
	if err != nil {
		return err
	} else if len(body) == 0 {
		return errors.New(string(body))
	}
	err = json.Unmarshal(body, response)

	if strings.Contains(string(body),"error") {
		jpushError:=&JpushError{}
		err=json.Unmarshal(body,jpushError)
		if err!=nil {
			return err
		}
		return errors.New(strconv.Itoa(jpushError.Error.Code))
	}
	return nil
}

func (p push) PushByTag(tag, title, context string, extra model.Extras) error {
	pushRequest := model.PushRequest{
		Platform: "all",
		Audience: model.Audience{
			Tag: []string{tag},
		},
		Notification: model.Notification{
			Android: model.Android{
				Alert:   title,
				Title:   title,
				BigText: context,
				Extras:  extra,
			},
			Ios: model.Ios{
				Alert:  title,
				Extras: extra,
			},
		},
	}
	return ResultSet(p.PostClient(consts.PushURL).JSON(pushRequest).Send())
}
