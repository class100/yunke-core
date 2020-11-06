package core

const (
	// FileApiUploadGet 获取文件上传信息
	FileApiUploadGet ApiPath = "courses/resources/uploads/infos"
	// FileApiDownloadGet 获取文件下载,打开信息
	FileApiDownloadGet ApiPath = "courses/resources/{fileId}/downloads"
	// FileApiDelete 根据文件编号删除
	FileApiDelete ApiPath = "courses/resources/{fileId}"
)
