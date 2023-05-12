package config

import (
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// NewConfig for application
// conf should be a pointer
func NewConfig(dirname string, envPref string, conf interface{}) {
	destV := reflect.ValueOf(conf)
	if destV.Kind() != reflect.Ptr && destV.Kind() != reflect.Interface {
		panic("conf should be a pointer or an interface")
	}
	config := viper.NewWithOptions(viper.KeyDelimiter("."))

	config.SetConfigName("application")
	config.SetConfigType("yaml")
	config.SetEnvPrefix(envPref)
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	config.AddConfigPath(dirname)
	config.AutomaticEnv()

	err := config.ReadInConfig()

	if err != nil {
		panic(err)
	}

	bindEnvs(config, conf)
	err = config.Unmarshal(conf)
	if err != nil {
		panic(err)
	}
}

func bindEnvs(c *viper.Viper, conf interface{}, parts ...string) {
	confV := reflect.ValueOf(conf)
	if confV.Kind() == reflect.Ptr || confV.Kind() == reflect.Interface {
		confV = confV.Elem()
	}
	confT := confV.Type()

	for i := 0; i < confT.NumField(); i++ {
		v := confV.Field(i)
		t := confT.Field(i)
		tag, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			tag = strings.ToLower(t.Name)
		}
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(c, v.Interface(), append(parts, tag)...)
		default:
			err := c.BindEnv(strings.Join(append(parts, tag), "."))
			if err != nil {
				panic(err)
			}
		}
	}

}
