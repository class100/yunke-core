package core

type (
	// JoinMeetingReq 加入会议请求
	JoinMeetingReq struct {
		MeetingData

		// AppId 产品ID
		AppId int64 `json:"appId,string"`
		// Username 用户名字
		Username string `json:"username"`
		// TODO ①以后删除掉
		// RoleType 角色类型
		// 0:默认
		// 1:督课员
		RoleType int8 `json:"roleType" validate:"omitempty,oneof=0 1"`
	}

	// JoinMeetingRsp 加入会议响应
	JoinMeetingRsp struct {
		// AppId 程序编号
		AppId int64 `json:"appId,string"`
		// VirtualUserId 虚拟账号
		VirtualUserId string `json:"virtualUserId"`
		// MeetingToken 令牌
		MeetingToken string `json:"meetingToken"`
		// MeetingPhone 会议中手机号
		MeetingPhone string `json:"meetingPhone"`
		// MeetingId 会议Id
		MeetingId string `json:"meetingId"`
		// MeetingNo 会议编号
		MeetingNo uint64 `json:"meetingNo,string"`
		// GroupId 组Id
		GroupId int64 `json:"groupId,string"`
		// GroupWebsocket 组websocket地址
		GroupWebsocket string `json:"groupWebsocket"`
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
