/*
 * BCS
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
type BatchAddPeersToChannelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddPeersToChannelResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchAddPeersToChannelResponse", string(data)}, " ")
}
