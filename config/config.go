package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// type Config struct {
// 	conf *viper.Viper
// }

var (
	viperConfMap map[string]*viper.Viper // 用来存放配置文件缓存
)

//读取顺序：环境变量>文件内容
func New(file string) error {
	//判断文件是否存在
	if _, err := os.Stat(file); err != nil {
		return err
	}
	//读取文件
	if err := readConfig(file); err != nil {
		return err
	}
	return nil
}

//读取文件内容
func readConfig(file string) error {
	bts, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if !in(strings.Split(path.Ext(file), ".")[1], viper.SupportedExts) {
		return fmt.Errorf("支持的文件类型：【%s\n】", viper.SupportedExts)
	}
	//fmt.Printf("配置文件类型：【%s】", path.Ext(file))
	v := viper.New()
	//设置文件类型
	v.SetConfigFile(path.Ext(file))
	//监听文件变化
	v.WatchConfig()
	if err = v.ReadConfig(bytes.NewBuffer(bts)); err != nil {
		return err
	}
	//环境变量替换
	if err = overrideEnv(v); err != nil {
		return err
	}
	if viperConfMap == nil {
		viperConfMap = make(map[string]*viper.Viper)
	}
	filename := strings.TrimSuffix(path.Base(file), path.Ext(file))
	viperConfMap[filename] = v
	return nil
}

//获取字符串配置信息
func GetStringConf(key string) string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return ""
	}
	v, ok := viperConfMap[keys[0]]
	if !ok {
		return ""
	}
	confString := v.GetString(strings.Join(keys[1:len(keys)], "."))
	return confString
}

//获取列表配置信息
func GetStringMapConf(key string) map[string]interface{} {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := viperConfMap[keys[0]]
	conf := v.GetStringMap(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func GetConf(key string) interface{} {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := viperConfMap[keys[0]]
	conf := v.Get(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取布尔值配置信息
func GetBoolConf(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return false
	}
	v := viperConfMap[keys[0]]
	conf := v.GetBool(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取64位浮点数配置信息
func GetFloat64Conf(key string) float64 {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := viperConfMap[keys[0]]
	conf := v.GetFloat64(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取Int类型配置信息
func GetIntConf(key string) int {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := viperConfMap[keys[0]]
	conf := v.GetInt(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取键值对配置信息
func GetStringMapStringConf(key string) map[string]string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := viperConfMap[keys[0]]
	conf := v.GetStringMapString(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取数组类型配置信息
func GetStringSliceConf(key string) []string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := viperConfMap[keys[0]]
	conf := v.GetStringSlice(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取时间类型配置信息
func GetTimeConf(key string) time.Time {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return time.Now()
	}
	v := viperConfMap[keys[0]]
	conf := v.GetTime(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取时间阶段配置信息
func GetDurationConf(key string) time.Duration {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := viperConfMap[keys[0]]
	conf := v.GetDuration(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//是否设置了key
func IsSetConf(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return false
	}
	v := viperConfMap[keys[0]]
	conf := v.IsSet(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//判断字符是否存在列表中
func in(str string, str_arr []string) bool {
	sort.Strings(str_arr)
	index := sort.SearchStrings(str_arr, str)
	if index < len(str_arr) && str_arr[index] == str {
		return true
	}
	return false
}
