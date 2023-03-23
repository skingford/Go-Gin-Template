/*
 * @Author: kingford
 * @Date: 2023-03-22 23:55:05
 * @LastEditTime: 2023-03-22 23:55:13
 */
// snowflake_id.go
package snowflakeid

import (
	"github.com/bwmarrin/snowflake"
)

func GenerateSnowflakeID(nodeID int64) (int64, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return 0, err
	}
	return node.Generate().Int64(), nil
}
