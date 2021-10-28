package main

import (
	"io"
	"os"
)

const (
	dll32    ="githubHelper_32.dll"
	dll64   ="githubHelper_64.dll"
	dllName = "sciter.dll"
)

func LoadDll()  {
	//先判断当前系统是否是64位系统
	bit := 32 << (^uint(0) >> 63)
	flag,err := PathExists("C:\\Windows\\System32\\sciter.dll")
	if err!= nil {
		panic(err)
	}
	//如果不存在
	if !flag {
		if bit==64 {
			os.Rename(dll64, dllName)
		}else {
			os.Rename(dll32, dllName)
		}
		_, err2 := CopyFile(dllName,"C:\\Windows\\System32\\sciter.dll")
		if err2 != nil {
			panic(err2)
		}
		os.Remove(dllName)
		os.Remove(dll32)
		os.Remove(dll64)
	}
}
//复制文件  需要考虑权限问题
func CopyFile(src, des string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		panic(err)
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}
//文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}