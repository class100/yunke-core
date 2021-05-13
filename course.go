package core

import `github.com/storezhang/gox`

const (
	// CourseTypeInteractiveClass 1：互动课
	CourseTypeInteractiveClass CourseType = 1
	// CourseTypeRecord 3：录播课
	CourseTypeRecord CourseType = 3
	// CourseTypePublic 4：公开课
	CourseTypePublic CourseType = 4

	// ClassModelCustom 1：定制的
	ClassModelCustom ClassMode = 1
	// ClassModelOriginal 1：原生模式
	ClassModelOriginal ClassMode = 2
)

type (
	// CourseType 课程类型
	// 1:互动课
	// 3：录播课
	// 4:公开课
	CourseType int8

	// ClassMode 上课模式
	ClassMode int8

	// Course 课程
	Course struct {
		gox.BaseStruct `xorm:"extends"`
		BaseSchool     `xorm:"extends"`

		// Name 课程名称
		Name string `xorm:"varchar(32) notnull default('')" json:"name" validate:"required,min=1,max=32"`
		// Type 课程类型
		// 1：小班课
		// 3：录播课
		// 4：公开课
		Type CourseType `default:"1" xorm:"tinyint notnull default(1)" json:"type" validate:"required"`
		// Cover 封面
		Cover string `xorm:"char(20) notnull default('')" json:"cover" validate:"omitempty,len=20"`
		// CreatorId 创建人
		CreatorId int64 `xorm:"bigint(20) notnull default(1)" json:"creatorId,string" validate:"required"`
		// Info 介绍
		Info string `xorm:"text(10000) notnull default('')" json:"info" validate:"omitempty,max=10000"`
		// ResourcePath 课程资源关联路径
		// 例如/a/b/c
		ResourcePath string `xorm:"varchar(255)" json:"resourcePath"`
		// ClassMode 上课模式
		// 1：自定义类型
		// 2：原生类型
		ClassMode ClassMode `xorm:"tinyint default(1)" json:"classMode"`
	}

	// AddCourseReq 添加课程
	AddCourseReq struct {
		// Name 课程名称
		Name string `json:"name" validate:"required,without_special_symbol,min=2,max=30"`
		// Type 课程类型
		// 1：小班课
		// 3：录播课
		// 4：公开课
		Type CourseType `json:"type" validate:"required_with=1 2 3 4"`
		// Cover 封面
		Cover string `json:"cover" validate:"omitempty,len=20"`
		// CreatorId 创建人
		CreatorId int64 `json:"creatorId,string" validate:"required"`
		// Info 介绍
		Info string `json:"info" validate:"omitempty,max=10000"`
		// ResourcePath 课程资源关联路径
		// 例如 /a/b/c
		ResourcePath string `json:"resourcePath" validate:"omitempty,startswith=/"`
		// ClassMode 上课模式
		// 1：自定义类型
		// 2：原生类型
		ClassMode ClassMode `default:"1" json:"classMode" validate:"omitempty,oneof=1 2"`
	}

	// UpdateCourseReq 更新课程请求
	UpdateCourseReq struct {
		gox.IdStruct

		// Name 课程名称
		Name string `json:"name" validate:"required,without_special_symbol,min=2,max=30"`
		// Type 课程类型
		// 1 小班课
		// 2 一对一
		// 3 录播课
		// 4 公开课
		Type CourseType `json:"type"`
		// CategoryId 课程分类
		CategoryId int64 `json:"categoryId,string"`
		// Cover 封面
		Cover string `json:"cover" validate:"omitempty,len=20"`
		// Info 介绍
		Info string `json:"info" validate:"omitempty,max=10000"`
		// ResourcePath 课程资源关联路径
		// 例如/a/b/c
		ResourcePath string `json:"resourcePath" validate:"omitempty,startswith=/"`
		// ClassMode 上课模式
		// 2：原生类型
		// 1：自定义类型
		ClassMode ClassMode `default:"1" json:"classMode" validate:"omitempty,oneof=1 2"`
	}
)
