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
type ShowUserMfaDeviceResponse struct {
	VirtualMfaDevice *MfaDeviceResult `json:"virtual_mfa_device,omitempty"`
}

func (o ShowUserMfaDeviceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowUserMfaDeviceResponse", string(data)}, " ")
}