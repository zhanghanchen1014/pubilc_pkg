package pkg

import (
	"context"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// 上传minio
func MUpload(filename, dst string) string {
	ctx := context.Background()
	endpoint := ""        //Minio 服务的地址(域名/IP+端日)
	accessKeyID := ""     //访问密钥(用户名)
	secretAccessKey := "" //访问密钥(密码)
	useSSL := false       //是否使用https

	// 创建minio客户端对象，用户后端操作
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err) // 初始化失败直接终止程序，输出错误
		return ""
	}

	// Make a new bucket called testbucket.
	bucketName := "2301atext" //自己设计的桶的名称(类似文件夹名，在Minio中需全局一)
	location := filename      //存储桶所在的“区域”(HinIO 逻辑，类似 AWSS3 的 Regon，这里用了入参 filename，可能需根据实际场景调整)

	//调用 MinI0 客户端的 MakeBucket 方法，尝试创建存储桶
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// 检查存储桶是否已存在
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists != nil || !exists {
			// 如果检查存在出错，或者桶确实不存在，才认为是错误
			log.Fatalln("创建存储桶失败且桶不存在:", err)
			return ""
		}
		// 如果桶已经存在，什么都不用做，继续执行上传逻辑
		log.Printf("存储桶 %s 已存在，继续上传...\n", bucketName)
	} else {
		log.Printf("Successfully created %s\n", bucketName) // 存储桶创建成功，打印提示
	}

	//1.定义上传参数
	objectName := filename //上传到 Minio 后，文件的对象名称
	filePath := dst        //本地文件路径
	// 根据文件扩展名设置更准确的 ContentType
	contentType := "application/octet-stream" // 默认类型
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".gif":
		contentType = "image/gif"
	case ".pdf":
		contentType = "application/pdf"
	case ".txt":
		contentType = "text/plain"
		// ... 可以添加更多类型
	}

	//2.执行上传:调用FPutObject 方法，将本地文件上传到Minio 存储桶
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	//3.上传成功后，打印反馈:输出上传的对象名、文件大小
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	//4.获取文件地址
	url, err := minioClient.PresignedGetObject(ctx, bucketName, objectName, time.Hour*24, nil)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return url.String()

}
