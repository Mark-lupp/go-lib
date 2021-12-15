package viper

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	FILE_TYPE = map[int]string{
		0: "yaml",
		1: "yml",
		2: "json",
		3: "toml",
	} // 配置文件后缀名格式
	ViperConfMap map[string]*viper.Viper // 用来存放配置文件缓存
)

const (
	YAML = iota
	YML
	JSON
	TOML
)

type Conf struct {
	Path     string
	fileType string
}

/*
path 配置文件路径
file_type 文件的后缀名支持yaml，yml，json，toml，默认yaml
*/
func New(path string, file_type ...int) (*Conf, error) {
	ft := ""
	switch file_type[0] {
	case YAML:
		ft = FILE_TYPE[YAML]
	case YML:
		ft = FILE_TYPE[YML]
	case JSON:
		ft = FILE_TYPE[JSON]
	case TOML:
		ft = FILE_TYPE[TOML]
	default:
		ft = FILE_TYPE[YAML]
	}
	conf := &Conf{
		Path:     path,
		fileType: ft,
	}
	if err := initViperConf(conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// Viper初始化配置文件
func initViperConf(c *Conf) error {
	f, err := os.Open(c.Path)
	if err != nil {
		return err
	}
	fileList, err := f.Readdir(1024)
	if err != nil {
		return err
	}
	for _, f0 := range fileList {
		if !f0.IsDir() {
			bts, err := ioutil.ReadFile(c.Path + "/" + f0.Name())
			if err != nil {
				return err
			}
			v := viper.New()
			v.SetConfigType(c.fileType)
			v.ReadConfig(bytes.NewBuffer(bts))
			pathArr := strings.Split(f0.Name(), ".")
			if ViperConfMap == nil {
				ViperConfMap = make(map[string]*viper.Viper)
			}
			ViperConfMap[pathArr[0]] = v
		}
	}
	return nil
}

//获取get配置信息
func (c *Conf) GetStringConf(key string) string {
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
func (c *Conf) GetStringMapConf(key string) map[string]interface{} {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetStringMap(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *Conf) GetConf(key string) interface{} {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := ViperConfMap[keys[0]]
	conf := v.Get(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *Conf) GetBoolConf(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return false
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetBool(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *Conf) GetFloat64Conf(key string) float64 {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetFloat64(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *Conf) GetIntConf(key string) int {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetInt(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *Conf) GetStringMapStringConf(key string) map[string]string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetStringMapString(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *Conf) GetStringSliceConf(key string) []string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetStringSlice(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取get配置信息
func (c *Conf) GetTimeConf(key string) time.Time {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return time.Now()
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetTime(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//获取时间阶段长度
func (c *Conf) GetDurationConf(key string) time.Duration {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetDuration(strings.Join(keys[1:len(keys)], "."))
	return conf
}

//是否设置了key
func (c *Conf) IsSetConf(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return false
	}
	v := ViperConfMap[keys[0]]
	conf := v.IsSet(strings.Join(keys[1:len(keys)], "."))
	return conf
}
