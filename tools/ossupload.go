package tools

import (
	"fmt"
	"mime/multipart"
	"os"
)

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func GetPlayUrl(filename string, content multipart.File) string {
	bucket := OssUpload()
	playurl := "Video address/" + filename
	// 将视频传到对应的文件夹
	err := bucket.PutObject(playurl, content)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return "https://yygh-lamo.oss-cn-beijing.aliyuncs.com/" + playurl
}

func GetCoverUrl(filename string) string {
	bucket := OssUpload()
	style := "video/snapshot,t_50000,f_jpg,w_800,h_600"
	// 指定视频路径
	ossVideoName := "Video address/" + filename
	// 生成带签名的URL，并指定过期时间为600s。
	signedURL, err := bucket.SignURL(ossVideoName, oss.HTTPGet, 2592000, oss.Process(style))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	return signedURL
}
