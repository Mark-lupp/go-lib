package viper

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/Mark-lupp/go-lib/dir"

	"github.com/spf13/viper"
)

var (
	ErrEmptyName = errors.New("文件名称不能为空！")
	ErrEmpty     = errors.New("文件不可为空！")
	ErrExistFile = errors.New("文件不存在！")
	ErrEmptyFile = errors.New("文件不可为空,请指定需读取的文件信息！")
)
var (
	ViperConfMap map[string]*viper.Viper // 用来存放配置文件缓存
)

// 读取配置
type ViperConfig struct {
	Debug     bool                     // 是否开发调试模式（控制台输出日志）
	FilePaths []string                 // 文件列表
	Watch     func(*viper.Viper) error // 监听配置文件变化
}

func New(vc ViperConfig) (config *ViperConfig, err error) {
	if len(vc.FilePaths) == 0 {
		return nil, ErrEmpty
	}
	for _, filepath := range vc.FilePaths {
		// 文件存在读取
		if dir.IsExist(filepath) {
			if err = readConfig(filepath, vc.Debug); err != nil {
				return
			}
		}
	}

	return nil, nil
}
func readConfig(path string, debug bool) error {
	_, err := os.Open(path)
	if err != nil {
		return err
	}
	v := viper.New()
	// 判断是否是文件
	if dir.IsFile(path) {
		bts, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		if debug {
			fmt.Printf("配置文件类型：【%s】", dir.ExtType(path))
		}
		v.SetConfigType(dir.ExtType(path))
		if err = v.ReadConfig(bytes.NewBuffer(bts)); err != nil {
			return err
		}
		if ViperConfMap == nil {
			ViperConfMap = make(map[string]*viper.Viper)
		}
		ViperConfMap[dir.Basename(path)] = v
	}
	return nil
}

// // Viper初始化配置文件
// func initViperConf(c *ViperConfig) error {
// 	f, err := os.Open(c.Path)
// 	if err != nil {
// 		return err
// 	}
// 	fileList, err := f.Readdir(1024)
// 	if err != nil {
// 		return err
// 	}
// 	for _, f0 := range fileList {
// 		if !f0.IsDir() {
// 			bts, err := ioutil.ReadFile(c.Path + "/" + f0.Name())
// 			if err != nil {
// 				return err
// 			}
// 			v := viper.New()
// 			v.SetConfigType(c.fileType)
// 			v.ReadConfig(bytes.NewBuffer(bts))
// 			pathArr := strings.Split(f0.Name(), ".")
// 			if ViperConfMap == nil {
// 				ViperConfMap = make(map[string]*viper.Viper)
// 			}
// 			ViperConfMap[pathArr[0]] = v
// 		}
// 	}
// 	return nil
// }

//获取get配置信息
func (c *ViperConfig) GetStringConf(key string) string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return ""
	}
	v, ok := ViperConfMap[keys[0]]
	if !ok {
		return ""
	}
	confString := v.GetString(strings.Join(keys[1:len(keys)], "."))
	return confString
}

//获取get配置信息
func (c *ViperConfig) GetStringMapConf(key string) map[string]interface{} {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetStringMap(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *ViperConfig) GetConf(key string) interface{} {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := ViperConfMap[keys[0]]
	conf := v.Get(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *ViperConfig) GetBoolConf(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return false
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetBool(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *ViperConfig) GetFloat64Conf(key string) float64 {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetFloat64(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *ViperConfig) GetIntConf(key string) int {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetInt(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *ViperConfig) GetStringMapStringConf(key string) map[string]string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetStringMapString(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *ViperConfig) GetStringSliceConf(key string) []string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetStringSlice(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *ViperConfig) GetTimeConf(key string) time.Time {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return time.Now()
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetTime(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取时间阶段长度
func (c *ViperConfig) GetDurationConf(key string) time.Duration {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetDuration(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//是否设置了key
func (c *ViperConfig) IsSetConf(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return false
	}
	v := ViperConfMap[keys[0]]
	conf := v.IsSet(strings.Join(keys[1:len(keys)], "."))
	return conf
}
