/*
 * @Author: kingford
 * @Date: 2023-03-23 16:32:33
 * @LastEditTime: 2023-03-23 17:07:07
 */
package crypto

import "github.com/gin-gonic/gin"

func (h *handler) Sign(c *gin.Context) {
	h.logger.Info("sign method logger")

	c.JSON(200, gin.H{
		"message": "Sign Method",
	})
}
