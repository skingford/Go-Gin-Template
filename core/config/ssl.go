/*
 * @Author: kingford
 * @Date: 2023-03-22 17:32:39
 * @LastEditTime: 2023-03-22 17:50:02
 */
package config

type Ssl struct {
	KeyStr string
	Pem    string
	Enable bool
	Domain string
}
