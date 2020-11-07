package core

const (
	// LectureApiAdd 添加课程章节
	LectureApiAdd ApiPath = "lectures"
	// LectureApiDeleteById 删除课程章节
	LectureApiDeleteById ApiPath = "lectures/{id}"
	// LectureApiUpdate 更新课程章节
	LectureApiUpdate ApiPath = "lectures/{id}"
	// LectureApiGetById 根据Id获取章节
	LectureApiGetById ApiPath = "lectures/{id}"
	// LectureApiGetByCourseId 根据课程Id获取章节信息
	LectureApiGetByCourseId ApiPath = "lectures/courses/{id}"
	// LectureSwitchSequence 调整章节顺序
	LectureSwitchSequence ApiPath = "lectures/s witches"
	// GetFirstLectureByCourse 获取第一
	LectureFirstByCourseId ApiPath = "lectures/firsts"
)
