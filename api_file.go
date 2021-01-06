package core

const (
	// FileApiUploadGet 获取文件上传信息
	FileApiUploadGet ApiPath = "files/uploads/infos"
	// FileApiDownloadGet 获取文件下载,打开信息
	FileApiDownloadGet ApiPath = "files/{fileId}/downloads/infos"
	// FileApiDelete 根据文件编号删除
	FileApiDelete ApiPath = "files/{fileId}"
	// FileApiCompletedUploadNotice 完成上传通知服务器上传成功
	FileApiCompletedUploadNotice = "files/{fileId}/notifies"
)
