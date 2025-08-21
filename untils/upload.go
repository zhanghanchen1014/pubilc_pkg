package untils

import (
	"errors"
	"github.com/zhanghanchen1014/pubilc_pkg/pkg"
)

// 定义统一上传方法
type Uploader interface {
	Upload(filename, dst string) error
	GetFileURL(filepath string) (string, error) // 新增方法：获取文件URL
}

// 阿里云上传
type ALiYunOss struct {
}

// 实现Uploader接口
func (u *ALiYunOss) Upload(filename, dst string) error {
	oss := pkg.OssUpload(filename, dst)
	if !oss {
		return errors.New("aliyun upload failed")
	}
	return nil
}

// 实现获取文件URL方法
func (u *ALiYunOss) GetFileURL(filepath string) (string, error) {
	url := pkg.GetOssUrl(filepath)
	if url == "" {
		return "", errors.New("生成OSS URL失败")
	}
	return url, nil
}

// Minio上传
type MinioOss struct {
}

// 实现Uploader接口
func (u *MinioOss) Upload(filename, dst string) error {
	minio := pkg.MUpload(filename, dst)
	if minio == "" || minio == "error" {
		return errors.New("minio upload failed")
	}
	return nil
}

// 实现获取文件URL方法
func (u *MinioOss) GetFileURL(filepath string) (string, error) {
	// 直接调用pkg.MUpload，因为它已经返回URL
	url := pkg.MUpload(filepath, filepath) // 注意：这里需要调整参数，或者最好重构pkg.MUpload
	if url == "" {
		return "", errors.New("生成Minio URL失败")
	}
	return url, nil
}

// 上传工厂
type UploadFactory struct{}

// 创建上传实例
func (f *UploadFactory) CreateUploader(uploadType string) (Uploader, error) {
	switch uploadType {
	case "aliyun":
		return &ALiYunOss{}, nil
	case "minio":
		return &MinioOss{}, nil
	default:
		return nil, errors.New("upload type not found")
	}
}
