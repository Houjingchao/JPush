package consts

//版本号
const (
	VersionV3 = "/v3"
)

//host
const (
	DeviceHost = "https://device.jpush.cn"
	PushHost   = "https://api.jpush.cn"
)

//device
const (
	/**
	查询设备的别名和标签
	-----分割线-----
	设置设备的别名与标签：
	tags: 支持add, remove 或者空字符串。当tags参数为空字符串的时候，表示清空所有的 tags；
	add/remove 下是增加或删除指定的 tag；
    一次 add/remove tag 的上限均为 100 个，且总长度均不能超过 1000 字节。
    可以多次调用 API 设置，一个注册 id tag 上限为1000个，应用 tag 总数没有限制 。
    alias: 更新设备的别名属性；当别名为空串时，删除指定设备的别名；
    mobile: 设备关联的手机号码
	 */
	DevicesURL = DeviceHost + VersionV3 + "/devices/"

	AliasURL = DeviceHost + VersionV3 + "/aliases/"

	TagsURL = DeviceHost + VersionV3 + "/tags/"
)

/* push url */
const (
	PushCidURL = PushHost + VersionV3 + "/push/cid"
	PushURL    = PushHost + VersionV3 + "/push"
)
