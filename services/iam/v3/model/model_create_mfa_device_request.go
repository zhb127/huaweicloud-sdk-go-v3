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

// Request Object
type CreateMfaDeviceRequest struct {
	Body *CreateMfaDeviceReq `json:"body,omitempty"`
}

func (o CreateMfaDeviceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateMfaDeviceRequest", string(data)}, " ")
}
