/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowQuotasResponse struct {
	Quotas         *Quotas `json:"quotas,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowQuotasResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowQuotasResponse", string(data)}, " ")
}
