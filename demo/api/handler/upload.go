package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanghanchen1014/pubilc_pkg/demo/api/resp"
	"github.com/zhanghanchen1014/pubilc_pkg/untils"
)

// oss上传
func AliyunUpload(c *gin.Context) {
	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		resp.Error(c, 400, "文件上传失败", err.Error())
		return
	}

	dst := "D:\\gowork\\src\\pubilc_pkg\\static\\upload\\" + file.Filename
	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		resp.Error(c, 400, "文件上传至本地失败", err.Error())
		return
	}

	// oss上传
	oss := untils.UploadFactory{}
	upload, err := oss.CreateUploader("aliyun")
	if err != nil {
		resp.Error(c, 400, "文件上传至oss失败", err.Error())
		return
	}
	err = upload.Upload(file.Filename, dst)
	if err != nil {
		resp.Error(c, 400, "文件上传至oss失败", err.Error())
		return
	}

	//获取文件的url
	url, err := upload.GetFileURL(file.Filename)
	if err != nil {
		resp.Error(c, 400, "获取文件URL失败", err.Error())
		return
	}

	resp.Success(c, "文件上传成功", url)
}

// minio上传
func MinioUpload(c *gin.Context) {
	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		resp.Error(c, 400, "文件上传失败", err.Error())
		return
	}

	dst := "D:\\gowork\\src\\pubilc_pkg\\static\\upload\\" + file.Filename
	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		resp.Error(c, 400, "文件上传至本地失败", err.Error())
		return
	}

	// minio上传
	minio := untils.UploadFactory{}
	upload, err := minio.CreateUploader("minio")
	if err != nil {
		resp.Error(c, 400, "文件上传至minio失败", err.Error())
		return
	}

	err = upload.Upload(file.Filename, dst)
	if err != nil {
		resp.Error(c, 400, "文件上传至minio失败", err.Error())
		return
	}

	//获取文件的url
	url, err := upload.GetFileURL(dst)
	if err != nil {
		resp.Error(c, 400, "获取文件URL失败", err.Error())
		return
	}

	resp.Success(c, "文件上传成功", url)
}
