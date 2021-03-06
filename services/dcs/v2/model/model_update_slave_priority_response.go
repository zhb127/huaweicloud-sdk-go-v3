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
type UpdateSlavePriorityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSlavePriorityResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSlavePriorityResponse", string(data)}, " ")
}
