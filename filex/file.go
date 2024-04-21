/*
 * @Author: Coller
 * @Date: 2021-09-24 12:30:08
 * @LastEditTime: 2024-01-04 09:30:25
 * @Desc: 文件操作
 */
package filex

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
)

/**
 * @desc: GetSize get the file size
 * @param undefined
 * @return {*}
 */
func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)

	return len(content), err
}

/**
 * @desc:  GetExt get the file ext
 * @param undefined
 * @return {*}
 */
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

/**
 * @desc: check if the file exists
 * @param undefined
 * @return {*}
 */
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

/**
 * @desc: check if the file has permission
 * @param undefined
 * @return {*}
 */
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

/**
 * @desc: create a directory if it does not exist
 * @return {*}
 */
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

/**
 * @desc: MkDir create a directory
 * @param undefined
 * @return {*}
 */
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

/**
 * @desc:  Open a file according to a specific mode
 * @param undefined
 * @param undefined
 * @param undefined
 * @return {*}
 */
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

/**
 * @desc: 文件目录是否存在
 * @return {*}
 */
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

/**
 * @desc: 批量创建文件夹
 * @param undefined
 * @return {*}
 */
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				fmt.Printf("create directory %v %v", v, err)
			}
		}
	}
	return err
}
