package oss

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/cool/coolfile"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

var (
	ctx          g.Ctx
	ossDriverObj = New()
)

type Oss struct {
	Client *oss.Client
	Bucket *oss.Bucket
}

func (m *Oss) New() coolfile.Driver {
	return m
}

func (m *Oss) GetMode() (data interface{}, err error) {
	data = g.MapStrStr{
		"mode": "local",
		"type": "oss",
	}
	return
}

func (m *Oss) Upload(ctx g.Ctx) (string, error) {
	var (
		err     error
		Request = g.RequestFromCtx(ctx)
	)

	file := Request.GetUploadFile("file")
	if file == nil {
		return "", gerror.New("上传文件为空")
	}

	src, err := file.Open()
	if err != nil {
		g.Log().Error(ctx, "文件打开失败")
	}
	defer src.Close()

	// 以当前年月日为目录
	dir := gtime.Now().Format("Ymd")
	fileName := Request.Get("key", grand.S(16, false)).String()
	fullPath := fmt.Sprintf("uploads/%s/%s", dir, fileName)

	// 创建目录
	err = m.Bucket.PutObject(fullPath, src)

	if err != nil {
		return "上传失败", err
	}

	url := fmt.Sprintf("https://%s.%s/%s", m.Bucket.BucketName, cool.Config.File.Oss.Endpoint, fullPath)

	return url, nil
}

func New() coolfile.Driver {
	ctx := context.Background()
	if cool.Config.File.Mode != "oss" {
		return nil
	}
	endpoint := cool.Config.File.Oss.Endpoint
	accessKeyID := cool.Config.File.Oss.AccessKeyID
	secretAccessKey := cool.Config.File.Oss.SecretAccessKey
	bucketName := cool.Config.File.Oss.BucketName
	// Initialize oss client object.
	client, err := oss.New(endpoint, accessKeyID, secretAccessKey)
	if err != nil {
		g.Log().Fatal(ctx, err)
		return nil
	}

	exist, err := client.IsBucketExist(bucketName)

	if err != nil {
		g.Log().Fatal(ctx, err)
		return nil
	}

	if exist {
		g.Log().Debug(ctx, fmt.Sprintf("存储桶%s已存在", bucketName))
	} else {
		// 创建存储桶
		err = client.CreateBucket(bucketName)
		if err != nil {
			g.Log().Fatal(ctx, err)
			return nil
		}
		g.Log().Debug(ctx, fmt.Sprintf("存储桶%s创建成功", bucketName))
	}

	bucket, _ := client.Bucket(bucketName)

	return &Oss{Client: client, Bucket: bucket}
}

func init() {
	var (
		err         error
		driverNames = g.SliceStr{"oss"}
	)

	if err != nil {
		panic(err)
	}

	for _, driverName := range driverNames {
		if err = coolfile.Register(driverName, ossDriverObj); err != nil {
			panic(err)
		}
	}
}
