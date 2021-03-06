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
type CreateTemporaryAccessKeyByTokenResponse struct {
	Credential     *Credential `json:"credential,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateTemporaryAccessKeyByTokenResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateTemporaryAccessKeyByTokenResponse", string(data)}, " ")
}
