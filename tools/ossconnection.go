package tools

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

// 上传文件到阿里云服务器
func OssUpload() *oss.Bucket {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-beijing.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := "LTAI5tPav8GX92F8CRD1F8Hm"
	accessKeySecret := "42jV6s92jYEynubnGcdB7G4Ma1ns5m"
	bucketName := "yygh-lamo"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}

	return bucket
}
