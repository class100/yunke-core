package core

import (
	`github.com/storezhang/gox`
)

type (
	// BatchAddCourseTimeReq 添加courseTimes请求
	BatchAddCourseTimeReq struct {
		CourseTimes []*CourseTime `json:"courseTimes"`
	}

	// BatchAddCourseTimeReq 添加courseTimes请求
	BatchDeleteCourseTimeReq struct {
		Ids gox.Int64Slice `json:"Ids"`
	}


	// CourseTime 课程时刻信息
	CourseTime struct {
		gox.IdStruct

		// CourseId 课程编号
		CourseId int64 `json:"courseId,string"`
		// TeacherId 主讲老师
		TeacherId int64 `json:"teacherId,string"`
		// ClassTime 上课时间
		ClassTime gox.Timestamp `json:"classTime"`
		// Duration 授课时长
		// 单位：分钟
		Duration int `json:"duration"`
		// ClassMode 上课模式
		// 0：原生类型
		// 1：自定义类型
		ClassMode ClassMode `default:"0" json:"classMode" validate:"omitempty,oneof=0 1"`
	}
)