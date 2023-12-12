package bootstrap

import (
	"algorithmplatform/global"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitializeConfig() *viper.Viper {
	var config string

	env := os.Getenv("GO_ENV")

	switch env {
	case "produce":
		config = "config-produce.yaml"
	default:
		config = "config.yaml"
	}
	fmt.Printf("config: %v\n", config)
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		//err =
		panic(fmt.Errorf("read config failed: %s ", err))
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
