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

// Response Object
type RestoreInstanceResponse struct {
	// 恢复记录ID。
	RestoreId      *string `json:"restore_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestoreInstanceResponse", string(data)}, " ")
}
