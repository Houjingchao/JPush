package DeviceAPI

type pushplatform int

const (
	PUSH_IOS     pushplatform = iota
	PUSH_ANDROID
	PUSH_WIN
	PUSH_ALL
)

var pushplatformString = [...]string{
	PUSH_IOS:     "ios",
	PUSH_ANDROID: "android",
	PUSH_WIN:     "winphone",
	PUSH_ALL:     "all",
}

func (p pushplatform) String() string {
	return pushplatformString[p]
}
