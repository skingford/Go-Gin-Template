/*
 * @Author: kingford
 * @Date: 2023-03-13 11:39:16
 * @LastEditTime: 2023-03-23 10:57:43
 */
package main

import (
	"go-gin-template/cmd"
)

func main() {
	// logger := pkg.NewZap()
	// defer logger.Sync()
	// logger.Info("==== Starting======")
	cmd.Execute()
}
