package pkg

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
		Mode string
	}
	Database struct {
		Host     string
		Port     int
		Username string
		Password string
	}
}

func NewViper() {
	// 初始化viper对象
	v := viper.New()

	// 获取当前工作目录的绝对路径
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("Failed to get working directory: %v", err))
	}

	// 设置viper配置
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(dir + "/config")
	v.AutomaticEnv()

	// 打印当前工作目录的绝对路径
	fmt.Printf("Current working directory: %s\n", dir+"/config")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to read config: %v", err))
	}

	// 映射配置到结构体
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %v", err))
	}

	// 打印配置
	fmt.Printf("Server config: %+v\n", config.Server)
	fmt.Printf("Database config: %+v\n", config.Database)
}
