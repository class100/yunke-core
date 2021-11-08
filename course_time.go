package core

import (
	`github.com/storezhang/gox`
)

const (
	// CourseTimeUnStart 0：未开始
	CourseTimeUnStart CourseTimeStatus = 0
	// CourseTimePrior 1: 预课中
	CourseTimePrior CourseTimeStatus = 1
	// CourseTimeStarting 2：进行中
	CourseTimeStarting CourseTimeStatus = 2
	// CourseTimeEnd 3：结束
	CourseTimeEnd CourseTimeStatus = 3
	// CourseTimeDelayed 4：拖堂
	CourseTimeDelayed CourseTimeStatus = 4
	// CourseTimeSuspended 5： 已经中止的
	CourseTimeSuspended CourseTimeStatus = 5

	// AllowedTypeStudentInCourseTime 1：课程时刻内的学生
	AllowedTypeStudentInCourseTime AllowedType = 1
	// AllowedTypeStudentInSys 2：系统内学生
	AllowedTypeStudentInSys AllowedType = 2
	// AllowedTypeTourist 3：允许游客
	AllowedTypeTourist AllowedType = 3

	// RecordTypeNo 1:不进行录制
	RecordTypeNo RecordType = 1
	// RecordTypeOnly 2:仅仅进行录制
	RecordTypeOnly RecordType = 2
	// RecordTypePlayback  3：录制并可以回放
	RecordTypePlayback RecordType = 3

	// 1 不允许直播
	LiveAllowedTypeNo LiveAllowedType = 1
	// 2 许直播
	LiveAllowedTypeYes LiveAllowedType = 2

	// 1-标清：不高于640 × 480
	VideoDimensionTypeStandard VideoDimensionType = 1
	// 2-高清：640 × 480 - 1280 × 720
	VideoDimensionTypeHigh VideoDimensionType = 1
	// 3-超清：高于1280 × 720
	VideoDimensionTypeSuper VideoDimensionType = 1
)

type (
	// 视频质量
	VideoDimensionType int8

	// BatchAddCourseTimeReq 批量添加课程时刻请求
	BatchAddCourseTimeReq struct {
		CourseTimes []*CourseTime `json:"courseTimes"`
	}

	// BatchDeleteCourseTimeReq 批量删除课程时刻请求
	BatchDeleteCourseTimeReq struct {
		Ids gox.Int64Slice `json:"Ids"`
	}

	// CourseTimeStatus 课程状态
	CourseTimeStatus int8

	// RecordType 课程计划录制操作类型
	// 1:不进行录制
	// 2:仅仅进行录制
	// 3：录制并可以回放
	RecordType int8

	// AllowedType 课程计划允许类型允许类型
	// 1：课程时刻内的学生
	// 2：系统内学生
	// 3：允许游客
	AllowedType int8

	// 互动课是否允许直播
	LiveAllowedType int8

	// CourseTime  课程时刻
	CourseTime struct {
		gox.IdStruct `xorm:"extends"`
		BaseSchool   `xorm:"extends"`

		// CourseScheduleId 课程计划
		CourseScheduleId int64 `xorm:"bigint(20) notnull default(1)" json:"courseScheduleId,string"`
		// CourseId 课程编号
		CourseId int64 `xorm:"bigint(20) default(1)" json:"courseId,string"`
		// TeacherId 主讲老师
		TeacherId int64 `xorm:"bigint(20) default(1)" json:"teacherId,string"`
		// 课节名称
		Name string `xorm:"varchar(255) default('课节')" json:"name,string"`
		// ClassTime 上课时间
		ClassTime gox.Timestamp `xorm:"datetime default('2020-02-04 09:55:52')" json:"classTime"`
		// Duration 授课时长
		// 单位：分钟
		Duration int `xorm:"int default(0)" json:"duration"`
		// Status 状态
		Status CourseTimeStatus `xorm:"int default(0)" json:"status"`
		// MeetingId 云视讯会议Id
		MeetingId string `xorm:"varchar(32) notnull default('')" json:"meetingId" validate:"required,min=1,max=32"`
		// MeetingNo 云视讯会议No
		MeetingNo uint64 `xorm:"bigint(20) notnull default(1)" json:"meetingNo,string"`
		// GroupId 组ID
		GroupId int64 `xorm:"bigint(20) notnull default(0)" json:"groupId,string"`
		// ClassMode 上课模式
		// 1：智慧课堂
		// 2：双师模式
		// 3：会议模式
		ClassMode ClassMode `xorm:"tinyint default(1)" json:"classMode"`
		// 允许类型
		// 1：课程时刻内的学生
		// 2：系统内学生
		// 3：允许游客
		AllowedType AllowedType `xorm:"tinyint default(1)" json:"allowedType"`
		// IsCaptchaRequired 是否需要验证码
		IsCaptchaRequired bool `xorm:"-" json:"-"`
		// Captcha 验证码
		Captcha *string `xorm:"char(6) default('')" json:"captcha"`
		// RecordType 录制操作类型
		// 1:不进行录制
		// 2:仅仅进行录制
		// 3：录制并可以回放
		RecordType RecordType `xorm:"tinyint default(1)" json:"recordType"`
		// 是否允许直播
		// 1 不允许直播
		// 2 允许直播
		LiveAllowedType LiveAllowedType `xorm:"tinyint default(1)" json:"liveAllowedType"`
		// 视频质量
		// 1-标清：不高于640 × 480
		// 2-高清：640 × 480 - 1280 × 720
		// 3-超清：高于1280 × 720
		VideoDimension VideoDimensionType `xorm:"tinyint default(1)"  json:"videoDimension"`
		// 上台人数
		StageNumber int8 `xorm:"tinyint default(1)"  json:"stageNumber"`
	}
)

func (bdctr *BatchDeleteCourseTimeReq) Models() (courseTimes []*CourseTime) {
	courseTimes = make([]*CourseTime, 0, len(bdctr.Ids))

	for _, id := range bdctr.Ids {
		courseTimes = append(courseTimes, &CourseTime{IdStruct: gox.IdStruct{Id: id}})
	}

	return
}

func (ct *CourseTime) GetEndTime() int64 {
	return ct.ClassTime.Time().Unix() + int64(ct.Duration+30)*60
}
