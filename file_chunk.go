package core

type (
	// InitiateMultipartUploadReq 初始化分块上传请求
	InitiateMultipartUploadReq struct {
		MultipartUploadData

		// UserId 请求者编号
		CreatorId int64 `json:"creatorId,string" validate:"required"`
	}

	// InitiateMultipartUploadRsp 初始化分块上传响应
	InitiateMultipartUploadRsp struct {
		// UploadId 腾讯Cos对象存储分块上传时唯一标识Id
		UploadId string `json:"uploadId"`
	}

	// CompleteMultipartUploadReq 完成分块上传请求
	CompleteMultipartUploadReq struct {
		MultipartUploadData

		// UploadId 腾讯Cos对象存储分块上传时唯一标识Id
		UploadId string `json:"uploadId" validate:"required"`
		// Parts 每一部分上传返回信息
		Parts []*MultipartUploadPartInfo `json:"parts" validate:"required,dive"`
	}

	// AbortMultipartUpload 中断分块上传请求
	AbortMultipartUploadReq struct {
		MultipartUploadData

		// UploadId 腾讯Cos对象存储分块上传时唯一标识Id
		UploadId string `json:"uploadId" validate:"required"`
	}

	// MultipartUploadData 分块上传数据
	MultipartUploadData struct {
		// FileId 文件编号
		FileId string `json:"fileId" validate:"required,len=20"`
		// Type 文件所属类型
		// 公共云盘:"public"
		// 私人云盘:"private"
		// 普通文件:"resource"
		// 课程文件："course"
		Type DirType `json:"type" validate:"required,oneof=course public private resource"`
	}

	// MultipartUploadPartInfo 分块上传块信息
	MultipartUploadPartInfo struct {
		// PartNumber 分块序列号
		PartNumber int `json:"partNumber" validate:"required"`
		// Etag 分块上传返回校验码
		Etag string `json:"etag" validate:"required"`
	}
)
