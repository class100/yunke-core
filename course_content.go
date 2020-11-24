package core

import (
	`github.com/storezhang/gox`
)

const (
	// ContentTypePicture 1:课程内容类型图片
	ContentTypePicture ContentType = 1
	// ContentTypeVideo 2:课程内容类型视频
	ContentTypeVideo ContentType = 2

	// ContentBelongTypeCourse 0：内容属于课程
	ContentBelongTypeCourse ContentBelongType = 0
	// ContentBelongTypeLecture 1：内容属于课程章节
	ContentBelongTypeLecture ContentBelongType = 1
)

type (
	// ContentType 课程内容类型
	ContentType int8

	// ContentBelongType 课程属于类型
	ContentBelongType int8

	// CourseContent 磁盘
	CourseContent struct {
		gox.BaseStruct `xorm:"extends"`

		// CourseId 课程编号
		CourseId int64 `xorm:"bigint(20) notnull default(1)" json:"courseId,string" validate:"required"`
		// 文件编号
		FileId string `xorm:"char(20) default('')" json:"fileId" validate:"required,len=20"`
		// FileName 文件名字
		FileName string `xorm:"varchar(255) default('')" json:"fileName" validate:"omitempty"`
		// Type 内容类型
		// 1:图片
		// 2:视频
		Type ContentType `xorm:"tinyint default(1)" json:"type"`
		// Sequence 顺序
		Sequence int `xorm:"int default(0)" json:"sequence"`
		// 创建者Id
		CreatorId int64 `xorm:"bigint(20) default(0)" json:"creatorId,string"`
	}

	// AddCourseContentReq 添加课程内容请求
	AddCourseContentReq struct {
		// Contents 课程内容,展示图片或者展示视频
		Contents []*CourseContent `json:"contents" validate:"dive"`
	}

	// AddCourseContentRsp 添加课程内容响应
	AddCourseContentRsp struct {
		// Contents 课程内容,展示图片或者展示视频
		Contents []*CourseContent `json:"contents"`
	}

	// GetCourseContentReq 获取课程内容请求
	GetCourseContentReq struct {
		// CourseId 课程编号
		CourseId int64 `json:"courseId,string" validate:"required"`
	}

	// GetCourseContentRsp 获取课程内容响应
	GetCourseContentRsp struct {
		// Contents 课程内容,展示图片或者展示视频
		Contents []*CourseContent `json:"contents"`
	}

	// DeleteCourseContentReq 删除课程内容请求
	DeleteCourseContentReq struct {
		// Ids 根据编号删除
		Ids gox.Int64Slice `json:"ids" validate:"omitempty"`
		// CourseId 课程编号
		CourseId int64 `json:"courseId,string" validate:"omitempty"`
	}
)
