package core

const (
	// UserApiAdd 添加用户
	UserApiAdd ApiPath = "users"
	// UserTeacherApiDelete 删除角色为老师的用户
	UserTeacherApiDelete ApiPath = "users/teachers/{id}"
	// UserApiUpdate 更新用户
	UserApiUpdate ApiPath = "users/{id}"
	// UserApiGetByPhone 根据电话获取用户信息
	UserApiGetByPhone = "users/phones/{phone}/infos"
)
