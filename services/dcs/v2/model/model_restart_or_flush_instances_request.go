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
type RestartOrFlushInstancesRequest struct {
	Body *ChangeInstanceStatusBody `json:"body,omitempty"`
}

func (o RestartOrFlushInstancesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestartOrFlushInstancesRequest", string(data)}, " ")
}
