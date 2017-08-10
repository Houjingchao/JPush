package main

import (
	"github.com/Houjingchao/JPush/v3/DeviceAPI"
	"fmt"
	"encoding/json"
	"github.com/Houjingchao/JPush/v3/DeviceAPI/model"
)

func main() {
	jp := DeviceAPI.NewJPush("appkey", "serctect")
	// 查询设备的别名与标签
	resq, err := jp.GetDevices("regid")
	if err != nil {
		panic(err)
	}
	res, _ := json.Marshal(resq)
	fmt.Println()
	fmt.Println(string(res))
	//设置设备的别名与标签
	err = jp.SetDevices("regid", model.SetDeviceRequestBody{
		SetTages: model.SetTages{
			Add:    &[]string{"testTag"},
			Remove: nil,
		},
		Alias:  nil,
		Mobile: nil,
	})
	fmt.Println(err)
	// 查询标签列表
	_, _ = jp.GetTags()//okay
	// 更新标签
	_ = jp.SetTag("tag", []string{}, []string{})
	// 删除标签
	_ = jp.DeleteTag("tag", 0) //0 "ios", 1 "android",2 "android,ios"
}
