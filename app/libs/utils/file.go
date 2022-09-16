package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// 向文本文本中追加内容(文件地址，内容)，不存在则创建
func AppendFile(filePath, content string) error {
	f, e := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		return e
	}
	defer f.Close()
	_, e = io.WriteString(f, content)
	return e
}

// 判断文件是否存在
func IsExist(fileAddr string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(fileAddr)
	if err != nil {
		if os.IsExist(err) { // 根据错误类型进行判断
			Println(err)
			return true
		}
		return false
	}
	return true
}

// 判断目录是否存在
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		Println(err)
		return false
	}
	return s.IsDir()
}

// 写入文件 文件名，文件字节
func WriteFile(fileName string, content []byte) error {
	return os.WriteFile(fileName, content, 0644)
}

//读取文件 文件名/绝对路径
func ReadFile(name string) ([]byte,error) {
	contents, err := os.ReadFile(name)
	if err != nil {
		return nil,err
	}
	return contents,nil
}

//zip 压缩 目录，压缩后的文件路径
func ZipDir(dir, zipFile string) error {
	fz, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer fz.Close()
	w := zip.NewWriter(fz)
	defer w.Close()
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fDest, err := w.Create(path[len(dir):])
			if err != nil {
				return err
			}
			fSrc, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fSrc.Close()
			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// 解压缩 zip 文件路径，解压缩之后的路径
func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			err := os.MkdirAll(fpath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}