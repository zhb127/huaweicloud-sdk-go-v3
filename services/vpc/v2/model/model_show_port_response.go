/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPortResponse struct {
	Port           *Port `json:"port,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowPortResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPortResponse", string(data)}, " ")
}
