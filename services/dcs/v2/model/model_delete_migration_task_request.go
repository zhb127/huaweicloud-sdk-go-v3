/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMigrationTaskRequest struct {
	Body *DeleteMigrateTaskRequest `json:"body,omitempty"`
}

func (o DeleteMigrationTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteMigrationTaskRequest", string(data)}, " ")
}
