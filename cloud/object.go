package cloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosClient struct {
	SecretID  string
	SecretKey string
	Region    string
}

func (c CosClient) CreatePresignedURL(bucket, path string) string {
	return CreatePresignedURL(c.SecretID, c.SecretKey, c.Region, bucket, path)
}

func CreatePresignedURL(serctID, sercetKey, region, bucket, path string) string {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket, region))
	su, _ := url.Parse(fmt.Sprintf("https://cos.%s.myqcloud.com", region))
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  serctID,   // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: sercetKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	presignedURL, err := client.Object.GetPresignedURL(context.Background(), http.MethodPut, path, serctID, sercetKey, time.Hour, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(presignedURL)
	return presignedURL.RawPath
}
