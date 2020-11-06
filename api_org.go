package core

const (
	// OrgApiClientPackageNotifyUrl 客户端打包通知地址
	OrgApiClientPackageNotifyUrl ApiPath = "clients/{id}/packages/notifies"
	// OrgApiClientAddUrl 添加客户端地址
	OrgApiClientAddUrl ApiPath = "clients"
	// OrgApiClientGetById 按编号获得客户端信息
	OrgApiClientGetById ApiPath = "clients/{id}"

	// OrgApiConfigUpdateUrl 更新配置地址
	OrgApiConfigUpdateUrl ApiPath = "configs/{name}"
)
