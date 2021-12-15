package db

import (
	"github.com/Mark-lupp/go-lib/lib/config"
	"github.com/Mark-lupp/go-lib/log"
	"github.com/minio/minio-go"
)

func initMinio() {
	var secure bool
	if config.GetMinioConfig().GetPath() == "s3.amazonaws.com" {
		secure = true
	}
	if minioClient, err = minio.New(config.GetMinioConfig().GetPath(), config.GetMinioConfig().GetAccessKeyId(), config.GetMinioConfig().GetSecretAccessKey(), secure); err != nil {
		panic(err)
	}
	log.NewLogger().Debug("minio success : " + config.GetMinioConfig().GetPath())
}
