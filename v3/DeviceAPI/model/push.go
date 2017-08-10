package model

//create by houjingchao on 17/08/
//定义  请求的结构体
type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	RegistrationId []string `json:"registration_id,omitempty"`
}

type Notification struct {
	Android struct {
		Alert     string `json:"alert"`
		Title     string `json:"title,omitempty"`
		BuilderID int `json:"builder_id,omitempty"`
		Extras struct {
			Newsid int `json:"newsid"`
		} `json:"extras,omitempty"`
	} `json:"android"`
	Ios struct {
		Alert string `json:"alert"`
		Sound string `json:"sound,omitempty"`
		Badge string `json:"badge,omitempty"`
		Extras struct {
			Newsid int `json:"newsid"`
		} `json:"extras,omitempty"`
	} `json:"ios"`
}

type Message struct {
	MsgContent  string `json:"msg_content"`
	ContentType string `json:"content_type"`
	Title       string `json:"title"`
	Extras struct {
		Key string `json:"key"`
	} `json:"extras"`
}

type SmsMessage struct {
	Content   string `json:"content"`
	DelayTime int `json:"delay_time"`
}

type Options struct {
	TimeToLive     int `json:"time_to_live"`
	ApnsProduction bool `json:"apns_production"`
	ApnsCollapseID string `json:"apns_collapse_id"`
}

type PushRequest struct {
	Cid          string `json:"cid,omitempty"`
	Platform     string `json:"platform"`
	Audience     Audience `json:"audience"`
	Notification Notification `json:"notification"`
	Message      Message  `json:"message,omitempty"`
	SmsMessage   SmsMessage `json:"sms_message,omitempty"`
	Options      Options  `json:"options,omitempty"`
}

/**
cid 的返回结果
 */

type CidResponse struct {
	Cidlist []string `json:"cidlist"`
}
