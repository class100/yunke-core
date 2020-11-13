package core

const (
	// FileChunkApiInitiateMultipartUpload 初始化分块上传信息
	FileChunkApiInitiateMultipartUpload ApiPath = "files/uploads/chunks"
	// FileChunkApiCompleteMultipartUpload 完成化分块上传信息
	FileChunkApiCompleteMultipartUpload ApiPath = "files/uploads/chunks"
	// FileApiChunkAbortMultipartUploadReq 取消分块上传
	FileChunkApiAbortMultipartUploadReq ApiPath = "files/uploads/chunks"
)
