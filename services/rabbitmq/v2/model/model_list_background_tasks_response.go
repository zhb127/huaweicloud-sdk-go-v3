/*
 * RabbitMQ
 *
 * RabbitMQ Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBackgroundTasksResponse struct {
	// 任务数量。
	TaskCount *string `json:"task_count,omitempty"`
	// 任务列表。
	Tasks          *[]ListBackgroundTasksRespTasks `json:"tasks,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListBackgroundTasksResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBackgroundTasksResponse", string(data)}, " ")
}
