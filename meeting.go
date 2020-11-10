package core

type (
	// JoinMeetingReq 加入会议请求
	JoinMeetingReq struct {
		MeetingData

		// AppId 产品ID
		AppId int64 `json:"appId,string"`
		// Username 用户名字
		Username string `json:"username"`
	}

	// JoinMeetingRsp 加入会议响应
	JoinMeetingRsp struct {
		// UserId 编号
		// 用户:则为UserId
		// 教室盒子：则为ClassRoomId
		UserID string `json:"userId,string"`
		// UserToken 令牌
		UserToken string `json:"userToken"`
		// VirtualPhone 虚拟手机号
		VirtualPhone string `json:"virtual_phone"`
		// MeetingId 会议Id
		MeetingId string `json:"meetingId"`
		// MeetingNo 会议编号
		MeetingNo uint64 `json:"meetingNo"`
		// GroupId 组Id
		GroupId int64 `json:"groupId,string"`
	}

	// LeaveMeetingReq 离开会议请求
	LeaveMeetingReq struct {
		MeetingData
	}

	// TerminateMeetingReq 终止会议请求
	TerminateMeetingReq struct {
		MeetingData
	}

	// 会议数据
	MeetingData struct {
		// UserId 编号
		// 用户:则为UserId
		// 教室盒子：则为ClassroomId
		UserId int64 `json:"userId,string" validate:"required"`
		// CourseTimeId 课时时刻编号
		CourseTimeId int64 `json:"courseTimeId,string" validate:"required"`
	}
)
