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
type ShowIpWhitelistRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowIpWhitelistRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowIpWhitelistRequest", string(data)}, " ")
}
