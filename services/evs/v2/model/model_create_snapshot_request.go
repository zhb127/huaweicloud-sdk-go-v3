/*
 * EVS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSnapshotRequest struct {
	Body *CreateSnapshotRequestBody `json:"body,omitempty"`
}

func (o CreateSnapshotRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSnapshotRequest", string(data)}, " ")
}
