package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestOverrideEnv(t *testing.T) {
	tests := []struct {
		name   string
		env    map[string]string
		file   string
		result map[string]interface{}
	}{
		{
			name: "no-env",
			env:  nil,
			file: "testdata/override-no-env.yaml",
			result: map[string]interface{}{
				"data1": "abc",
				"data2": 123,
				"data3": []interface{}{1, 2},
				"data4": map[string]interface{}{
					"a": 1,
					"b": 2,
				},
				"data5": []interface{}{
					map[string]interface{}{
						"a": 1,
					},
				},
			},
		},
		{
			name: "full-env",
			env: map[string]string{
				"TEST_1":           "DEF",
				"TEST_2":           "456",
				"TEST_3_1":         "2",
				"TEST_3_2_NOT_SET": "",
				"TEST_4_A":         "3",
				"TEST_5_A":         "3",
			},
			file: "testdata/override-env.yaml",
			result: map[string]interface{}{
				"data1": "DEF",
				"data2": "456",
				"data3": []interface{}{
					"2", "2",
				},
				"data4": map[string]interface{}{
					"a": "3",
					"b": "1",
				},
				"data5": []interface{}{
					map[string]interface{}{
						"a": "3",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// load env
			for k, v := range tt.env {
				os.Setenv(k, v)
			}

			// load config
			v := viper.New()
			v.SetConfigType("yaml")
			content, err := ioutil.ReadFile(tt.file)
			if err != nil {
				t.Fatalf("read config file error: %v", err)
			}
			if err := v.ReadConfig(bytes.NewReader(content)); err != nil {
				t.Fatalf("read config error: %v", err)
			}

			// environment override
			if err := overrideEnv(v); err != nil {
				t.Fatalf("override environment failure: %v", err)
			}

			// verify result
			realSettings := v.AllSettings()
			if !reflect.DeepEqual(realSettings, tt.result) {
				t.Fatalf("validate config failure: \n excepted: \n%v\n actual: \n%v", tt.result, realSettings)
			}
		})
	}
}
