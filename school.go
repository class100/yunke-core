package core

type (
	// SchoolId 学校编号
	BaseSchool struct {
		// SchoolId 学校编号
		SchoolId int64 `xorm:"bigint(20) notnull default(1)" json:"schoolId,string" validate:"required"`
	}
)
