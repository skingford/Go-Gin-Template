/*
 * @Author: kingford
 * @Date: 2023-03-11 00:57:58
 * @LastEditTime: 2023-03-11 01:06:48
 */
package pkg

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// export APP_ENV=prod
func NewViper() {
	// 设置配置文件的搜索路径
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	// 从环境变量中获取当前环境
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // 默认为开发环境
	}

	// 加载共享配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error loading config file:", err)
		return
	}

	// 读取特定环境的配置文件
	viper.SetConfigName(fmt.Sprintf("config_%s", env))
	err = viper.MergeInConfig()
	if err != nil {
		fmt.Println("Error loading config file:", err)
		return
	}

	// 获取配置选项
	port := viper.GetInt("http.port")
	dbHost := viper.GetString("db.host")
	dbPort := viper.GetInt("db.port")

	// 输出配置选项
	fmt.Printf("HTTP port: %d\n", port)
	fmt.Printf("DB host: %s\n", dbHost)
	fmt.Printf("DB port: %d\n", dbPort)
}
