package DeviceAPI

import (
	"testing"
	"github.com/Houjingchao/JPush/v3/DeviceAPI/model"
	"fmt"
)

func TestNewPush(t *testing.T) {
    //var b chan bool = make(chan bool)
	//for i:=0;i<30;i++{
		 pushss()
	//}
	//<-b
}

func pushss() (){
	jp := NewJPush("test", "test")
	extra:=model.Extras{
		Action:"action",
		Collect:"collect",
		Func:"func",
		Url:"url",
	}
	//for i:=0;i<=900;i++ {
		// 查询设备的别名与标签
		//fmt.Println(i)
		err := jp.PushByRegid("test", "test", "test", extra)
		fmt.Println(err)
	//}
}
