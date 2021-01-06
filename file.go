package core

const (
	// StorageStateEnough 0：表示空间充足
	StorageStateEnough StorageState = 0
	// StorageStateNotEnough 1：表示空间不足
	StorageStateNotEnough StorageState = 1

	// OperationTypeDownload 1:直接下载
	OperationTypeDownload OperationType = 1
	// OperationTypeOpen 2:打开
	OperationTypeOpen OperationType = 2
	// OperationTypeEdit 3:编辑
	OperationTypeEdit OperationType = 3
)

type (
	// OperationType 操作类型
	OperationType int8

	// StorageState 内存空间是否足够标记
	StorageState int8

	// BaseUploadReq 文件上传请求基础信息
	BaseUploadReq struct {
		// FileId 文件编号
		FileId *string `json:"fileId" validate:"omitempty,len=20"`
		// Filename 文件名字
		Filename string `json:"filename:" validate:"omitempty,without_special_symbol,filename"`
		// DataSize 文件大小
		// 单位B
		DataSize int64 `json:"dataSize" validate:"required"`
		// DirType 文件上传类型,如果不是功能特殊,基本上都用resource
		// 系统文件 system 如：用户头像，系统本身Log
		// 普通文件 resource 如：作业附件
		// 公共云盘 public
		// 私有云盘 private
		DirType DirType `json:"dirType" validate:"required,oneof=system resource public private"`
	}

	// BaseUploadRsp 文件上传响应基础信息
	BaseUploadRsp struct {
		BaseFileReq

		// Url 上传地址
		Url string `json:"url"`
		// State 云盘状态
		// 0：表示空间足够
		// 1：表示空间不足
		State StorageState `json:"state"`
	}

	// BaseDownloadReq 获取文件下载,打开,编辑基础请求
	BaseDownloadReq struct {
		BaseFileReq

		// Name 文件另存名字
		Name string `json:"name" validate:"omitempty,min=1,without_special_symbol,filename"`
		// Type 下载类型
		// 1 立即下载
		// 2 打开
		// 3 在线编辑
		Type OperationType `default:"1" json:"type" validate:"required,oneof=1 2 3"`
	}

	// BaseDownloadRsp 获取文件下载,打开,编辑响应
	BaseDownloadRsp struct {
		// Url 预签名Url
		Url string `json:"url"`
	}

	// BaseFileReq 基础文件信息
	BaseFileReq struct {
		// FileId 文件编号
		FileId string `json:"fileId" validate:"required,len=20"`
	}
)
