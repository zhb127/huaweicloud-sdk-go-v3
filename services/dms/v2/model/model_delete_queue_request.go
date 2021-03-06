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
type DeleteQueueRequest struct {
	ProjectId string `json:"project_id"`
	QueueId   string `json:"queue_id"`
}

func (o DeleteQueueRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteQueueRequest", string(data)}, " ")
}
