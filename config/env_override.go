package config

import (
	"os"
	"regexp"

	"github.com/spf13/viper"
)

const EnvRegularExpression = "\\${(?P<ENV>[_A-Z0-9]+):(?P<DEF>.*)}"

func overrideEnv(v *viper.Viper) error {
	envRegex, err := regexp.Compile(EnvRegularExpression)
	if err != nil {
		return err
	}

	keys := v.AllKeys()
	for _, key := range keys {
		overrideConfig(v, key, envRegex)
	}
	return nil
}

func overrideConfig(v *viper.Viper, key string, envRegex *regexp.Regexp) {
	confValue := v.Get(key)
	switch val := confValue.(type) {
	case string:
		v.Set(key, overrideString(val, envRegex))
	case []interface{}:
		v.Set(key, overrideSlice(val, envRegex))
	case int:
		v.Set(key, val)
	case bool:
		v.Set(key, val)
	}
}

func overrideSlice(val []interface{}, envRegex *regexp.Regexp) []interface{} {
	res := make([]interface{}, 0)
	for _, perValue := range val {
		switch v := perValue.(type) {
		case string:
			res = append(res, overrideString(v, envRegex))
		case map[interface{}]interface{}:
			res = append(res, overrideMapInterfaceInterface(v, envRegex))
		case map[string]interface{}:
			res = append(res, overrideMapStringInterface(v, envRegex))
		default:
			res = append(res, v)
		}
	}
	return res
}

func overrideString(val string, envRegex *regexp.Regexp) string {
	groups := envRegex.FindStringSubmatch(val)
	if len(groups) == 0 {
		return val
	}

	if v := os.Getenv(groups[1]); v != "" {
		return v
	}
	return groups[2]
}

func overrideMapInterfaceInterface(val map[interface{}]interface{}, regex *regexp.Regexp) interface{} {
	cfg := make(map[string]interface{})
	for k, v := range val {
		cfg[k.(string)] = v
	}
	return overrideMapStringInterface(cfg, regex)
}

func overrideMapStringInterface(val map[string]interface{}, regex *regexp.Regexp) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range val {
		switch d := v.(type) {
		case string:
			res[k] = overrideString(d, regex)
		case []interface{}:
			res[k] = overrideSlice(d, regex)
		case map[string]interface{}:
			res[k] = overrideMapStringInterface(d, regex)
		case map[interface{}]interface{}:
			res[k] = overrideMapInterfaceInterface(d, regex)
		default:
			res[k] = d
		}
	}
	return res
}
