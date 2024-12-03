package convert

import (
	"github.com/spf13/cast"
	serverv1 "kratos-gin-template/api/server/v1"
)

func PlatformString(platform serverv1.DevicePlatform) string {
	switch platform {
	case serverv1.DevicePlatform_DevicePlatformAndroid:
		return "android"
	case serverv1.DevicePlatform_DevicePlatformIOS:
		return "ios"
	case serverv1.DevicePlatform_DevicePlatformWeb:
		return "web"
	case serverv1.DevicePlatform_DevicePlatformWechatWeb:
		return "wechat_web"
	case serverv1.DevicePlatform_DevicePlatformWechatMiniApp:
		return "wechat_mini_app"
	case serverv1.DevicePlatform_DevicePlatformUndefined:
		fallthrough
	default:
		return "unknown"
	}
}

func PlatformPB(platform any) serverv1.DevicePlatform {
	switch p := platform.(type) {
	case serverv1.DevicePlatform:
		return p
	case int, int8, int32, int64:
		return serverv1.DevicePlatform(cast.ToInt(platform))
	case string:
		break
	default:
		return serverv1.DevicePlatform_DevicePlatformUndefined
	}

	switch platform {
	case "android":
		return serverv1.DevicePlatform_DevicePlatformAndroid
	case "ios":
		return serverv1.DevicePlatform_DevicePlatformIOS
	case "web":
		return serverv1.DevicePlatform_DevicePlatformWeb
	case "wechat_web":
		return serverv1.DevicePlatform_DevicePlatformWechatWeb
	case "wechat_mini_app":
		return serverv1.DevicePlatform_DevicePlatformWechatMiniApp
	default:
		return serverv1.DevicePlatform_DevicePlatformUndefined
	}
}
