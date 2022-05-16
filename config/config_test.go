package config

import "testing"

func TestNew(t *testing.T) {
	filePath := "testdata/config-env.yaml"
	if err := New(filePath); err != nil {
		t.Errorf("read config error: %v", err)
	}
	r := GetStringConf("config-env.testa.data")
	t.Logf("val %s\n", r)
}
