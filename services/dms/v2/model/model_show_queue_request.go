/*
 * DMS
 *
 * DMS Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowQueueRequest struct {
	ProjectId         string `json:"project_id"`
	QueueId           string `json:"queue_id"`
	IncludeDeadletter *bool  `json:"include_deadletter,omitempty"`
}

func (o ShowQueueRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowQueueRequest", string(data)}, " ")
}
