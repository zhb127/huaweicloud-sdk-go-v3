/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBackgroundTasksRequest struct {
	ProjectId  string  `json:"project_id"`
	InstanceId string  `json:"instance_id"`
	Start      *int32  `json:"start,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
	BeginTime  *string `json:"begin_time,omitempty"`
	EndTime    *string `json:"end_time,omitempty"`
}

func (o ListBackgroundTasksRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBackgroundTasksRequest", string(data)}, " ")
}
