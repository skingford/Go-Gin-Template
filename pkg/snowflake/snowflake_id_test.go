/*
 * @Author: kingford
 * @Date: 2023-03-22 23:55:25
 * @LastEditTime: 2023-03-23 21:06:39
 */
// snowflake_id_test.go
package snowflakeid

import (
	"testing"
)

func TestGenerateSnowflakeID(t *testing.T) {
	nodeID := int64(1)

	s, _ := NewSnowflakeIDGenerator(nodeID)
	id, err := s.Generate()

	if err != nil {
		t.Fatalf("Error generating snowflake ID: %v", err)
	}
	if id == 0 {
		t.Error("Expected snowflake ID not to be 0")
	}

	// Generate another ID to ensure uniqueness
	id2, err := s.Generate()

	if err != nil {
		t.Fatalf("Error generating another snowflake ID: %v", err)
	}
	if id2 == 0 {
		t.Error("Expected second snowflake ID not to be 0")
	}
	if id == id2 {
		t.Error("Expected snowflake IDs to be unique")
	}

	t.Logf("id1: %d", id)
	t.Logf("id2: %d", id2)
}
