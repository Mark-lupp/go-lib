package db

import (
	"fmt"
	"os"

	"github.com/Mark-lupp/go-lib/lib/config"
	"github.com/Mark-lupp/go-lib/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go.uber.org/zap"
)

//mysql连接池
func initMysql() {
	var err error
	sql := fmt.Sprintf("%s:%s@(%s:%d)/%s", config.GetMysqlConfig().GetUser(), config.GetMysqlConfig().GetPwd(),
		config.GetMysqlConfig().GetIp(), config.GetMysqlConfig().GetPort(), config.GetMysqlConfig().GetDbName())
	log.NewLogger().Debug("[initMysql] " + sql)
	mysqlEngine, err = xorm.NewEngine("mysql", sql)
	if err != nil {
		log.NewLogger().Error("[initMysql] "+sql, zap.Error(err))
		os.Exit(0)
	}
	mysqlEngine.SetMaxOpenConns(config.GetMysqlConfig().GetPoolSize())
	mysqlEngine.SetMaxIdleConns(config.GetMysqlConfig().GetPoolSize())
	if err = mysqlEngine.Ping(); err != nil {
		panic(err)
	}
}

func CloseMysqlConnection() {
	_ = mysqlEngine.Close()
}
