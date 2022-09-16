package upload

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
)

// 文件上传类
/**
 * @Example:
	var c *gin.Context
	file, err := c.FormFile("file")
	yun := upload.Upload{File:file,Format:"images"}
	fmt.Println(yun.Upload())
*/
type Upload struct {
	File   *multipart.FileHeader // 上传的文件
	Format string                // 上传的类型
	suffix string                // 文件后缀
}

// 上传文件，返回上传成功后的地址
func (up Upload) Upload() string {
	up.checkFormat()
	up.checkSuffix()
	//up.checkType()
	up.checkSize()
	return up.save()
}

func (up *Upload) checkFormat() {
	if format[up.Format] == "" {
		utils.ExitError("不允许的类型", -1)
	}
}

// 验证文件后缀
func (up *Upload) checkSuffix() {
	up.suffix = strings.ToLower(path.Ext(up.File.Filename)) // 取文件后缀
	up.suffix = strings.ToLower(up.suffix)
	if suffixList[up.Format][up.suffix] == "" {
		utils.ExitError("文件格式错误", -1)
	}
}

// 验证文件类型
func (up *Upload) checkType() {
	types := up.File.Header.Values("Content-Type")
	fmt.Println(up.File)
	if len(types) == 0 || formatList[up.Format][types[0]] == "" {
		utils.ExitError("文件类型错误", -1)
	}
}

// 验证文件大小
func (up *Upload) checkSize() {
	if up.File.Size > sizeList[up.Format]*1024*1024 {
		utils.ExitError("上传的图片过大", -1)
	}
}

// 保存文件
func (up *Upload) save() string {
	// 打开文件
	src, e := up.File.Open()
	file, e := up.File.Open()
	if e != nil {
		utils.ExitError(e.Error(), -1)
	}
	defer src.Close()
	name := up.getFileMd5(file) + up.suffix // 获取文件的唯一名称
	fileName := up.getFilePath() + name     // 保存文件的绝对路径
	if utils.IsExist(fileName) {            // 避免文件重复保存
		return name
	}
	out, e := os.Create(fileName)
	if e != nil {
		utils.ExitError(e.Error(), -1)
	}
	defer out.Close()
	_, e = io.Copy(out, src)
	if e != nil {
		utils.ExitError(e.Error(), -1)
	}
	return name
}

// 获取文件md5
func (up *Upload) getFileMd5(f multipart.File) string {
	md5hash := md5.New()
	_, _ = io.Copy(md5hash, f)
	return fmt.Sprintf("%x", md5hash.Sum(nil))
}

// 获取文件保存的绝对路径
func (up *Upload) getFilePath() string {
	dir := config.App.Other.PublicDir
	if dir == "" {
		utils.ExitError("未设置 public_dir 文件上传静态目录", -1)
	}
	if !utils.IsDir(dir) {
		_, err := os.Create(dir)
		if err != nil {
			utils.ExitError(fmt.Sprintf("创建目录失败,%v", err), -1)
		}
	}
	return dir
}
