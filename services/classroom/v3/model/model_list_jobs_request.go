/*
 * Classroom
 *
 * devcloud classedge api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListJobsRequest struct {
	SourceFrom string `json:"source_from"`
	SourceId   string `json:"source_id"`
	Offset     *int32 `json:"offset,omitempty"`
	Limit      *int32 `json:"limit,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
