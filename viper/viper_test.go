package viper

import (
	"testing"
)

func TestNew(t *testing.T) {
	config := &ViperConfig{
		Debug: true,
		FilePaths: []string{
			"testdata/conf-no-env.yaml",
		},
	}
	c, err := New(*config)
	if err != nil {
		t.Error(err)
	}
	c.GetStringConf("testa.data")
}
