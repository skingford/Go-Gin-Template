/*
 * @Author: kingford
 * @Date: 2023-03-21 17:07:30
 * @LastEditTime: 2023-03-21 17:07:33
 */
package config

type Application struct {
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          int64
	Name          string
	JwtSecret     string
	Mode          string
	DemoMsg       string
	EnableDP      bool
}

var ApplicationConfig = new(Application)
