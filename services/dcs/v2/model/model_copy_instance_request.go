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
type CopyInstanceRequest struct {
	InstanceId string              `json:"instance_id"`
	Body       *BackupInstanceBody `json:"body,omitempty"`
}

func (o CopyInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CopyInstanceRequest", string(data)}, " ")
}
