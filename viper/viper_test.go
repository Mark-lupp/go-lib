package viper

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	config := &ViperConfig{
		Debug: true,
		FilePaths: []string{
			"/conf/application.yaml",
		},
	}
	c, err := New(*config)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	c.GetConf("base.mysql")
}
