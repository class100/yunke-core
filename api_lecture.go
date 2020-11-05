package core

const (
	// 课程章节
	LectureApiAdd           ApiPath = "lectures"
	LectureApiDeleteById    ApiPath = "lectures/{id}"
	LectureApiUpdate        ApiPath = "lectures/{id}"
	LectureApiGetById       ApiPath = "lectures/{id}"
	LectureApiGetByCourseId ApiPath = "lectures/courses/{id}"
	LectureSwitchSequence   ApiPath = "lectures/switches"
)
