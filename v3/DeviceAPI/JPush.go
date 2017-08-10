package DeviceAPI

import b64 "encoding/base64"

type JPush struct {
	Authorization string
	Device
	Push//用于推送
}

func NewJPush(appKey, masterSecrect string) (*JPush) {
	authorization := "Basic " + b64.StdEncoding.EncodeToString([]byte(appKey+":"+masterSecrect))
	return &JPush{
		Authorization: authorization,
		Device:        NewDevice(authorization),
		Push:          NewPush(authorization),
	}
}
