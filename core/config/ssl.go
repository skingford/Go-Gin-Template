/*
 * @Author: kingford
 * @Date: 2023-03-21 17:56:15
 * @LastEditTime: 2023-03-21 17:56:18
 */
package config

type Ssl struct {
	KeyStr string
	Pem    string
	Enable bool
	Domain string
}

var SslConfig = new(Ssl)
