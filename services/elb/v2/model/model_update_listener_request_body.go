/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateListenerRequestBody struct {
	Listener *UpdateListenerReq `json:"listener"`
}

func (o UpdateListenerRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateListenerRequestBody", string(data)}, " ")
}
