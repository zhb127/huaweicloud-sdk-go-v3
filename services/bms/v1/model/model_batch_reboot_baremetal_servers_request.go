/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchRebootBaremetalServersRequest struct {
	Body *RebootBody `json:"body,omitempty"`
}

func (o BatchRebootBaremetalServersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchRebootBaremetalServersRequest", string(data)}, " ")
}
