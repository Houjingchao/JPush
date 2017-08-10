package DeviceAPI

import (
	"github.com/Houjingchao/JPush/v3/DeviceAPI/model"
	"github.com/cocotyty/httpclient"
	"github.com/Houjingchao/JPush/v3/DeviceAPI/consts"
)

type (
	Device interface {
		GetDevices(registrationID string) (model.Devices, error)                   // 查询设备的别名与标签
		SetDevices(registrationID string, option model.SetDeviceRequestBody) error // 设置设备的别名与标签
		GetTags() (model.Tags, error)                                              // 查询标签列表
		SetTag(tag string, add, remove []string) error                             // 更新标签
		DeleteTag(tag string, pl platform) error                                   //删除标签
	}
	device struct {
		Authorization string
	}
)

func NewDevice(authorization string) Device {
	return device{Authorization: authorization}
}

//client Authorization 已经初始化了
func (d device) PostClient(url, param string) *httpclient.HttpRequest {
	return httpclient.Post(url + param).Head("Authorization", d.Authorization)
}

func (d device) GetClient(url string, param string) *httpclient.HttpRequest {
	return httpclient.Get(url + param).Head("Authorization", d.Authorization)
}
func (d device) GetDevices(registrationID string) (model.Devices, error) {
	ds := model.Devices{}
	return ds, ResultGet(d.GetClient(consts.DevicesURL, registrationID).Send(), &ds)
}

func (d device) SetDevices(registrationID string, option model.SetDeviceRequestBody) error {
	return ResultSet(d.PostClient(consts.DevicesURL, registrationID).JSON(option).Send())
}

func (d device) GetTags() (model.Tags, error) { //获取所有的标签
	ts := model.Tags{}
	return ts, ResultGet(d.GetClient(consts.TagsURL, "").Send(), &ts)
}

//为一个标签添加或者删除设备。

func (d device) SetTag(tag string, add, remove []string) error {
	if len(add) == 0 {
		add = nil
	}
	if len(remove) == 0 {
		remove = nil
	}
	requePram := model.SetTagRequestBody{
		SetTages: model.SetTages{
			Add:    &add,
			Remove: &remove,
		},
	}
	return ResultSet(d.PostClient(consts.TagsURL, tag).JSON(requePram).Send())
}

//删除一个标签，以及标签与设备之间的关联关系
//platform 可选参数，不填则默认为所有平台。
func (d device) DeleteTag(tag string, pl platform) error {
	return ResultSet(d.TagDelete(tag, pl).Send())
}

func (d device) TagDelete(tag string, pl platform) *httpclient.HttpRequest {
	return httpclient.Delete(consts.TagsURL + tag + "?platform=" + pl.String()).Head("Authorization", d.Authorization)
}
