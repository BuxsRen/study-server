package upload

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
)

// 七牛云
/**
 * @Example:
	var c *gin.Context
	file, err := c.FormFile("file")
	yun := upload.QiNiuYun{File:file,Types:"images"}
	fmt.Println(yun.Upload())
*/
type QiNiuYun struct {
	File         *multipart.FileHeader // 上传的文件
	Format       string                // 上传的类型
	bucket       string
	access_key   string
	secret_key   string
	mac          *qbox.Mac
	token        string
	formUploader *storage.FormUploader
	ret          storage.PutRet
	putExtra     storage.PutExtra
}

var qiniuyun *QiNiuYun

func init() {
	yun := &QiNiuYun{}
	qiniuyun = yun.new()
}

// 上传文件到七牛云，返回上传成功后的地址
func (qny *QiNiuYun) Upload() string {
	f := Upload{File: qny.File, Format: qny.Format}
	filename := f.Upload()
	path := config.App.Other.PublicDir + filename
	err := qiniuyun.uploadToQiNIuYun(filename, path)
	if err != nil {
		utils.ExitError("上传失败", -1)
	}
	return filename
}

func (qny *QiNiuYun) new() *QiNiuYun {
	yun := QiNiuYun{
		bucket:     config.App.QiNiuYun.Bucket,
		access_key: config.App.QiNiuYun.AccessKey,
		secret_key: config.App.QiNiuYun.SecretKey,
	}
	if yun.bucket == "" || yun.access_key == "" || yun.secret_key == "" {
		utils.ExitError("请先配置七牛云设置", -1)
	}
	yun.mac = qbox.NewMac(yun.access_key, yun.secret_key)

	putPolicy := storage.PutPolicy{
		Scope: yun.bucket,
	}
	yun.token = putPolicy.UploadToken(yun.mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneXinjiapo
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	yun.formUploader = storage.NewFormUploader(&cfg)
	yun.ret = storage.PutRet{}
	// 可选配置
	yun.putExtra = storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	return &yun
}

// 上传文件 文件名(推荐md5后的唯一文件名:xxxx.txt) 文件绝对路径(/www/wwwwroot/public/xxx.txt)
func (qny *QiNiuYun) uploadToQiNIuYun(fileName, filePath string) error {
	return qny.formUploader.PutFile(context.Background(), &qny.ret, qny.token, fileName, filePath, &qny.putExtra)
}
