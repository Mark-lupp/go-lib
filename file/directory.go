package file

import (
	"fmt"
	"os"
)

//@function: PathExists
//@description: 文件目录是否存在
//@param: path string "文件目录路径"
//@return: bool, error "存在true不存在false，调用失败错误"
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

//@function: CreateDir
//@description: 批量创建文件夹
//@param: dirs ...string "文件夹路径数组"
//@return: err error "调用失败错误"
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				return fmt.Errorf("create directory: %v Fail,error: %v", v, err)
			}
		}
	}
	return err
}
