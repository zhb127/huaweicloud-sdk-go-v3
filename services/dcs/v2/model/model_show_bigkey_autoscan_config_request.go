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
type ShowBigkeyAutoscanConfigRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowBigkeyAutoscanConfigRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowBigkeyAutoscanConfigRequest", string(data)}, " ")
}
