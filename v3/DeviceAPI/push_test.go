package DeviceAPI

import (
	"testing"
	"github.com/Houjingchao/JPush/v3/DeviceAPI/model"
	"fmt"
)

func TestNewPush(t *testing.T) {
	jp := NewJPush("appkey", "masterSecrect")
	extra:=model.Extras{
		Action:"action",
		Collect:"collect",
		Func:"func",
		Url:"url",
	}
	// 查询设备的别名与标签
	err:=jp.PushByRegid("regid", "test", "test", extra)
	fmt.Println(err)
}
