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
type ShowInstanceRequest struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceRequest", string(data)}, " ")
}
