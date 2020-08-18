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
type CreatePermanentAccessKeyResponse struct {
	Credential *CreateCredentialResult `json:"credential,omitempty"`
}

func (o CreatePermanentAccessKeyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePermanentAccessKeyResponse", string(data)}, " ")
}