package yunke

const (
	// 版本号
	// ApiVersionDefault 无版本，默认
	ApiVersionDefault ApiVersion = "default"
	// ApiVersionV1 版本1
	ApiVersionV1 ApiVersion = "v1"

	// UrlApiPrefix Api前缀
	UrlApiPrefix string = "api"
)

type (
	// ApiVersion 版本
	ApiVersion string
)
