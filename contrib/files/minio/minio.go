package minio

import (
	"context"
	"fmt"
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/cool/coolfile"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	ctx            g.Ctx
	minioDriverObj = New()
)

type Minio struct {
	Client     *minio.Client
	BucketName string
}

func (m *Minio) New() coolfile.Driver {
	g.Log().Debug(ctx, m, m.BucketName)
	return m
}

func (m *Minio) GetMode() (data interface{}, err error) {
	data = g.MapStrStr{
		"mode": "local",
		"type": "minio",
	}
	return
}

func (m *Minio) Upload(ctx g.Ctx) (string, error) {
	g.Log().Debug(ctx, m)

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

	g.Log().Debug(ctx, fullPath)
	// 创建目录
	info, err := m.Client.PutObject(ctx, m.BucketName, fullPath, src, -1, minio.PutObjectOptions{})

	g.Log().Debug(ctx, info)
	if err != nil {
		return "上传失败", err
	}
	return info.Location, nil
}

func New() coolfile.Driver {
	ctx := context.Background()
	if cool.Config.File.Mode != "minio" {
		return nil
	}
	endpoint := cool.Config.File.Oss.Endpoint
	accessKeyID := cool.Config.File.Oss.AccessKeyID
	secretAccessKey := cool.Config.File.Oss.SecretAccessKey
	useSSL := cool.Config.File.Oss.UseSSL
	bucketName := cool.Config.File.Oss.BucketName
	location := cool.Config.File.Oss.Location
	// Initialize minio client object.
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		g.Log().Error(ctx, "初始化Minio失败")
		return nil
	}

	if client.IsOffline() {
		g.Log().Error(ctx, "Minio当前不在线")
		return nil
	}

	err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			g.Log().Debug(ctx, fmt.Sprintf("存储桶%s已存在", bucketName))
		} else {
			g.Log().Fatal(ctx, err)
			return nil
		}
	} else {
		g.Log().Info(ctx, fmt.Sprintf("存储桶%s创建成功", bucketName))
	}
	return &Minio{Client: client, BucketName: bucketName}
}

func init() {
	var (
		err         error
		driverNames = g.SliceStr{"minio"}
	)

	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	for _, driverName := range driverNames {
		if err = coolfile.Register(driverName, minioDriverObj); err != nil {
			panic(err)
		}
	}
}
