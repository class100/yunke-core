package core

import (
	`encoding/json`

	`github.com/storezhang/gox`
)

const (
	// ProductPC 包括windows和mac
	ProductPC Product = "pc"
	// ProductWeb 产品页面
	ProductWeb Product = "web"
	// ProductConsole 管理后台
	ProductConsole Product = "console"
	// ProductApp 包括Android和iOS
	ProductApp Product = "app"
)

const (
	// 默认值
	// DefaultPCStartupLogo PC端默认启动图标
	DefaultPCStartupLogo string = "pc-startup-logo.ico"
	// DefaultPCConfigFilename PC端默认配置文件名
	DefaultPCConfigFilename string = "resources/lib/conf/conf.json"
	// DefaultPCPackageConfigFilename PC端包配置文件名
	DefaultPCPackageConfigFilename string = "resources/app/package.json"
)

const (
	// PackageStatusPackaged 已打包
	PackageStatusPackaged PackageStatus = 1
	// PackageStatusNotPackage 未打包
	PackageStatusNotPackage PackageStatus = 2
)

type (
	// Product 产品
	Product string

	// PackageStatus 打包状态
	PackageStatus uint8

	// Custom 定制化
	Custom struct {
		// Product 产品
		Product Product `xorm:"pk varchar(16) notnull default('')" json:"product"`
		// Config 配置
		Config map[string]interface{} `xorm:"json default(null)" json:"config"`
		// UpdatedAt 更新时间
		UpdatedAt gox.Timestamp `xorm:"updated default('2020-06-11 09:55:52')" json:"updatedAt"`
	}
)

func (c Custom) String() string {
	jsonBytes, _ := json.MarshalIndent(c, "", "    ")

	return string(jsonBytes)
}
