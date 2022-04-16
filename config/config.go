package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var once sync.Once
var conf []Config

func GetConfig() ([]Config, error) {
	var err error
	once.Do(func() {
		viper.AddConfigPath(".")
		var homeDir string
		homeDir, err = os.UserHomeDir()
		if err != nil {
			err = fmt.Errorf("获取用户目录失败 %w", err)
			return
		}
		viper.AddConfigPath(filepath.Join(homeDir, ".proxy-router"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		// 读取配置文件
		err = viper.ReadInConfig()
		if err != nil {
			err = fmt.Errorf("读取配置文件失败 %w", err)
			return
		}
		serverConfigs := cast.ToSlice(viper.Get("servers"))
		if serverConfigs == nil || len(serverConfigs) == 0 {
			err = errors.New("没有配置服务")
			return
		}
		for i := range serverConfigs {
			var routers []RouterConfig
			for j := range cast.ToSlice(viper.Get(fmt.Sprintf("servers.%d.routes", i))) {
				confMap := viper.GetStringMap(fmt.Sprintf("servers.%d.routes.%d", i, j))
				var config RouterConfig
				switch confMap["type"] {
				case "root":
					config = NewRootRouterConfig(
						cast.ToString(confMap["prefix"]),
						cast.ToString(confMap["staticPath"]),
						cast.ToStringSlice(confMap["indexPath"])...)
				case "proxy":
					config = NewProxyRouterConfig(
						cast.ToString(confMap["prefix"]),
						cast.ToString(confMap["scheme"]),
						cast.ToString(confMap["host"]),
						cast.ToString(confMap["baseUri"]))
				}
				if cast.ToBool(confMap["core"]) {
					// config = NewCore
					config.EnableCore()
				}
				routers = append(routers, config)
			}
			conf = append(conf, NewConfig(
				viper.GetUint(fmt.Sprintf("servers.%d.port", i)),
				routers...))
		}
	})
	return conf, err
}
