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
type CreateMigrationTaskRequest struct {
	Body *CreateMigrationTaskBody `json:"body,omitempty"`
}

func (o CreateMigrationTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateMigrationTaskRequest", string(data)}, " ")
}
