/*
 * @Author: kingford
 * @Date: 2023-03-22 17:49:25
 * @LastEditTime: 2023-03-22 17:59:03
 */
package config

type Application struct {
	Server
	Ssl
	Database
}

var ApplicationConfig = new(Application)
