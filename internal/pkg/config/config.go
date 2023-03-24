package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-es/config"
)

var v *viper.Viper

func init() {
	v = viper.New()

	v.SetConfigType("yml")
	v.AddConfigPath("./config")
	v.SetEnvPrefix("go-es")
	v.AutomaticEnv()
}

func LoadConfig(env string) {
	configName := "settings.yml"
	if env != "" {
		configName = fmt.Sprintf("settings.%s.yml", env)
	}

	// 读取配置文件
	v.SetConfigName(configName)
	if err := v.ReadInConfig(); err != nil {
		panic("启动失败，err: 读取配置文件 " + configName + " 失败. " + err.Error())
	}

	// 加载配置
	if err := v.Unmarshal(config.GlobalConfig); err != nil {
		panic("启动失败，err: 加载配置失败，" + err.Error())
	}

	// 监控配置文件，变更时重新加载，无需重启
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		//  todo 需要记录日志
		if err := v.Unmarshal(config.GlobalConfig); err != nil {
			// todo 需要记录日志
		}
	})
}
