package pkg

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"log"
	"time"
)

// 简单上传
func OssUpload(filename, dst string) bool {
	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider("", "")).
		WithRegion("cn-shanghai") //写入自己的OSS地域

	// 创建OSS客户端
	client := oss.NewClient(cfg)

	// 填写要上传的本地文件路径和文件名称，例如 /Users/localpath/exampleobject.txt
	localFile := dst

	// 创建上传对象的请求
	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr("2301a"),         // 存储空间名称 桶
		Key:          oss.Ptr(filename),        // 对象名称
		StorageClass: oss.StorageClassStandard, // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPrivate,     // 指定对象的访问权限为私有访问
		Metadata: map[string]string{
			"yourMetadataKey1": "yourMetadataValue1", // 设置对象的元数据
		},
	}

	// 执行上传对象的请求
	result, err := client.PutObjectFromFile(context.TODO(), putRequest, localFile)
	if err != nil {
		log.Fatalf("failed to put object from file %v", err)
	}

	// 打印上传对象的结果
	log.Printf("put object from file result:%#v\n", result)
	return true
}

// 预签名下载
func GetOssUrl(filename string) string {
	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider("", "")).
		WithRegion("cn-shanghai") //写入自己的OSS地域

	// 创建OSS客户端
	client := oss.NewClient(cfg)

	// 生成GetObject的预签名URL
	result, err := client.Presign(context.TODO(), &oss.GetObjectRequest{
		Bucket: oss.Ptr("2301a"),
		Key:    oss.Ptr(filename),
	},
		oss.PresignExpires(10*time.Minute),
	)
	if err != nil {
		log.Fatalf("failed to get object presign %v", err)
		return ""
	}

	log.Printf("request method:%v\n", result.Method)
	log.Printf("request expiration:%v\n", result.Expiration)
	log.Printf("request url:%v\n", result.URL)

	return result.URL
}
