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
type CreateListenerRequestBody struct {
	Listener *CreateListenerReq `json:"listener"`
}

func (o CreateListenerRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateListenerRequestBody", string(data)}, " ")
}
