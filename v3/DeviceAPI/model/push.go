package model

//create by houjingchao on 17/08/
//定义  请求的结构体
type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	RegistrationId []string `json:"registration_id,omitempty"`
}
type Android struct {
	Alert     string `json:"alert"`
	Title     string `json:"title,omitempty"`
	BuilderID int `json:"builder_id,omitempty"`
	BigText   string `json:"big_text"`
	Extras    Extras `json:"extras,omitempty"`
}
type Notification struct {
	Android Android `json:"android"`
	Ios     Ios `json:"ios"`
}
type Ios struct {
	Alert          string `json:"alert"`
	Sound          string `json:"sound,omitempty"`
	Badge          string `json:"badge,omitempty"`
	//MutableContent string `json:"mutable-content"`
	Extras         Extras `json:"extras,omitempty"`
}
type Extras struct {
	//根据自己的业务添加的
	Action  string `json:"action"`
	Collect string `json:"collect"`
	Func    string `json:"func"`
	Url     string `json:"url"`
}
type Message struct {
	MsgContent  string `json:"msg_content"`
	ContentType string `json:"content_type"`
	Title       string `json:"title"`
	Extras      Extras `json:"extras,omitempty"`
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
	Message      Message  `json:"-"`
	SmsMessage   SmsMessage `json:"-"`
	Options      Options  `json:"-"`
}

/**
cid 的返回结果
 */

type CidResponse struct {
	Cidlist []string `json:"cidlist"`
}
