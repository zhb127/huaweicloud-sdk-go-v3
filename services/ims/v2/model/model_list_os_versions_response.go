/*
 * IMS
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
type ListOsVersionsResponse struct {
	Body []ListOsVersionsResponseBody `json:"body,omitempty"`
}

func (o ListOsVersionsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListOsVersionsResponse", string(data)}, " ")
}