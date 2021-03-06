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
type CreateHotkeyScanTaskRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o CreateHotkeyScanTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateHotkeyScanTaskRequest", string(data)}, " ")
}
