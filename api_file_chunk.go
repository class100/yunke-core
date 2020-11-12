package core

const (
	// FileChunkApiInitiateMultipartUpload 初始化分块上传信息
	FileChunkApiInitiateMultipartUpload ApiPath = "/uploads/multiparts/chunks"
	// FileChunkApiCompleteMultipartUpload 完成化分块上传信息
	FileChunkApiCompleteMultipartUpload ApiPath = "/uploads/multiparts/chunks"
	// FileApiChunkAbortMultipartUploadReq 取消分块上传
	FileChunkApiAbortMultipartUploadReq ApiPath = "/uploads/multiparts/chunks"
)
