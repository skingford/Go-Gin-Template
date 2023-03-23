/*
 * @Author: kingford
 * @Date: 2023-03-22 09:44:10
 * @LastEditTime: 2023-03-23 10:01:07
 */
package pkg

import (
	"fmt"
	"go-gin-template/common/config"
	"os"

	"github.com/spf13/viper"
)

func NewViper() {
	// 初始化viper对象
	v := viper.New()

	// 获取当前工作目录的绝对路径
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get working directory: %v", err))
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
		panic(fmt.Errorf("failed to read config: %v", err))
	}

	// 映射配置到结构体

	if err := v.Unmarshal(&config.ApplicationConfig); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %v", err))
	}

	// 打印配置
	fmt.Printf("Server config: %+v\n", config.ApplicationConfig.Server)
	fmt.Printf("Database config: %+v\n", config.ApplicationConfig.Database)
}
