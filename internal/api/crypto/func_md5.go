/*
 * @Author: kingford
 * @Date: 2023-03-23 16:32:23
 * @LastEditTime: 2023-03-23 17:06:13
 */
package crypto

import "github.com/gin-gonic/gin"

func (h *handler) Md5(c *gin.Context) {
	h.logger.Info("md5 method logger")

	c.JSON(200, gin.H{
		"message": "Md5 Method",
	})
}
