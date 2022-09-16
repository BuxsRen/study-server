package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 用于build时的解压缩

func main(){
	str := os.Args
	if len(str) > 2 {
		e := ZipDir(str[1], str[2])
		if e != nil {
			fmt.Println(e)
		}
	}
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