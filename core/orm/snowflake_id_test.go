/*
 * @Author: kingford
 * @Date: 2023-03-22 23:55:25
 * @LastEditTime: 2023-03-22 23:58:11
 */
// snowflake_id_test.go
package snowflakeid

import (
	"testing"
)

func TestGenerateSnowflakeID(t *testing.T) {
	nodeID := int64(1)
	id, err := GenerateSnowflakeID(nodeID)

	t.Fatal("id1:", id)

	if err != nil {
		t.Fatalf("Error generating snowflake ID: %v", err)
	}
	if id == 0 {
		t.Error("Expected snowflake ID not to be 0")
	}

	// Generate another ID to ensure uniqueness
	id2, err := GenerateSnowflakeID(nodeID)

	t.Fatal("id2:", id2)

	if err != nil {
		t.Fatalf("Error generating another snowflake ID: %v", err)
	}
	if id2 == 0 {
		t.Error("Expected second snowflake ID not to be 0")
	}
	if id == id2 {
		t.Error("Expected snowflake IDs to be unique")
	}
}
