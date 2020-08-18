/*
 * IAM
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
type ShowPermanentAccessKeyResponse struct {
	Credential *ShowCredential `json:"credential,omitempty"`
}

func (o ShowPermanentAccessKeyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPermanentAccessKeyResponse", string(data)}, " ")
}