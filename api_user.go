package core

const (
	// UserApiAdd 添加用户
	UserApiAdd ApiPath = "users"
	// UserTeacherApiDelete 删除角色为老师的用户
	UserTeacherApiDelete ApiPath = "users/teachers/{id}"
	// UserApiUpdate 更新用户
	UserApiUpdate ApiPath = "users/{id}"
	// UserApiGet 根据电话获取用户信息
	UserApiGet = "users/{id}"
	// UserApiBatchAdd 批量添加用户
	UserApiBatchAdd = "users/batches"
)
