package base

import (
	"github.com/Mark-lupp/go-lib/lib/config"
	"github.com/Mark-lupp/go-lib/lib/db"
)

//配置文件的目录
func Init(path string) {
	config.Init(path)
	db.Init()
}
