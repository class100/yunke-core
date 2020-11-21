package core

import (
	`github.com/storezhang/gox`
)

const (
	// UserStatusNormal 正常状态
	UserStatusNormal UserStatus = 0

	// GetUserTypeId 默认通过用户编号
	GetUserTypeId GetUserType = "id"
	// GetUserInfoTypePhone 通过手机号获取
	GetUserTypePhone GetUserType = "phone"
)

type (
	// GetUserType 获取用户信息类型
	GetUserType string

	// UserStatus 用户状态
	UserStatus int8

	// User 用户实体
	User struct {
		gox.BaseUser `xorm:"extends"`

		// Nickname 昵称
		Nickname string `xorm:"varchar(64) notnull default('')" json:"nickname"`
		// Avatar 头像
		Avatar string `xorm:"char(20) notnull default('')" json:"avatar"`
		// FirstName 姓
		FirstName string `xorm:"varchar(16) notnull default('')" json:"firstName"`
		// LastName 名
		LastName string `xorm:"varchar(16) notnull default('')" json:"lastName"`
		// Sex 性别
		// 0是女
		// 1是男
		// 是不是很形象
		Sex int8 `default:"1" xorm:"tinyint notnull index('idx_sex')  default(1)" json:"sex"`
		// Status 用户状态
		// 0-正常
		// 1-停用
		Status UserStatus `default:"0" xorm:"tinyint notnull default(0)" json:"status"`
		// Info 介绍
		Info string `xorm:"varchar(255) notnull default('')" json:"info"`
		// SchoolNum 学号
		SchoolNum gox.DBString `xorm:"varchar(64) notnull default('')" json:"schoolNum"`
	}

	// SimpleRole 简单角色数据
	SimpleRole struct {
		// Id 角色编号
		Id int64 `json:"id,string"`
		// Name 角色类型
		// 1:管理员
		// 2:讲师
		// 3:助教
		// 4:学生
		// 5:教务
		// 6:运营
		Name string `json:"name"`
	}

	// OmitemptyPhone 手机号（有数据就验证）
	OmitemptyPhone struct {
		// 手机号
		Phone string `json:"phone" validate:"omitempty,mobile"`
	}

	// OmitemptySchoolNum 学号（有数据就验证）
	OmitemptySchoolNum struct {
		// SchoolNum 学号
		SchoolNum string `json:"schoolNum" validate:"omitempty,without_special_symbol,min=2,max=30"`
	}

	// RequiredBasePhone 手机号（必须）
	RequiredBasePhone struct {
		// Phone 手机号
		// 类似于+86-17089792784
		// 所有的手机号都必须带区号
		// 默认是+86
		// 只能用-分隔
		Phone string `json:"phone" validate:"required,mobile"`
	}

	// AddUserReq 添加用户请求
	AddUserReq struct {
		RequiredBasePhone
		OmitemptySchoolNum

		// Username 用户名
		Username string `json:"username" validate:"omitempty,min=1,max=30,email"`
		// Nickname 昵称
		Nickname string `json:"nickname" validate:"required,without_special_symbol,min=2,max=30"`
		// Avatar 头像
		Avatar string `json:"avatar" validate:"omitempty,len=20"`
		// FirstName 姓
		FirstName string `json:"firstName" validate:"omitempty,min=1,max=16"`
		// LastName 名
		LastName string `json:"lastName" validate:"omitempty,min=1,max=16"`
		// Sex 性别
		// 0是女
		// 1是男
		// 是不是很形象
		Sex int8 `default:"1" json:"sex" validate:"omitempty,oneof=0 1"`
		// Info 介绍
		Info string `json:"info" validate:"omitempty,max=255"`
		// RoleNames 该用户所属于的角色名称列表
		RoleNames []string `json:"roleNames" validate:"required"`
		// ClassIds 所属的班级
		ClassIds gox.Int64Slice `json:"classIds"`
		// Status 用户状态
		Status UserStatus `default:"0" json:"status" validate:"omitempty,oneof=0 1"`
		// Password 密码
		Password string `json:"password" validate:"omitempty,min=8,max=30"`
	}

	// UpdateUserReq 更新用户请求
	UpdateUserReq struct {
		OmitemptyPhone
		OmitemptySchoolNum

		// Username 用户名
		Username string `json:"username" validate:"omitempty,min=1,max=30,email"`
		// Nickname 昵称
		Nickname string `json:"nickname" validate:"omitempty,without_special_symbol,min=2,max=30"`
		// Avatar 头像
		Avatar string `json:"avatar" validate:"omitempty,len=20"`
		// FirstName 姓
		FirstName string `json:"firstName" validate:"omitempty,min=1,max=16"`
		// LastName 名
		LastName string `json:"lastName" validate:"omitempty,min=1,max=16"`
		// Sex 性别
		// 0是女
		// 1是男
		// 是不是很形象
		Sex int8 `default:"1" json:"sex" validate:"omitempty,oneof=0 1"`
		// Info 介绍
		Info string `json:"info" validate:"omitempty,max=255"`
		// ClassIds 班级ID
		ClassIds gox.Int64Slice `json:"classIds"`
		// Status 用户状态
		Status UserStatus `default:"0" json:"status" validate:"omitempty,oneof=0 1"`
		// RoleNames 该用户所属于的角色名称列表
		RoleNames []string `json:"roleNames"`
	}

	//  GetUserReq 根据手机号获取用户信息请求
	GetUserReq struct {
		// Id 需要查询的用户的唯一标识
		Id string `json:"id" validate:"required"`
		// Type 取得用户的类型
		Type GetUserType `default:"id" json:"type" validate:"required,oneof=id phone email schoolNum"`
		// Extra 是否包含额外信息
		Extra bool `default:"true" json:"extra"`
	}

	GetUserRsp struct {
		User

		// Roles 角色列表
		Roles []*SimpleRole `json:"roleNames"`
	}
)
