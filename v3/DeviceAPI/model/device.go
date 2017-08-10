package model

type (
	/**
	 查询设备的别名与标签
	 {
     "tags": ["tag1", "tag2"],
     "alias": "alias1",
     "mobile": "13012345678"
     }
	 */
	Devices struct {
		//返回信息结构
		Tags   []string `json:"tags"`
		Alias  *string `json:"alias"`
		Mobile int `json:"mobile"`
	}

	SetDeviceRequestBody struct {// 请求request
		SetTages `json:"tags"`
		Alias  *string `json:"alias,omitempty"`//当别名为空串时，删除指定设备的别名； 故而选择了omitempty
		Mobile *string `json:"mobile,omitempty"`
	}

	SetTages struct {
		Add    *[]string `json:"add,omitempty"`
		Remove *[]string `json:"remove,omitempty"`
	}
)
