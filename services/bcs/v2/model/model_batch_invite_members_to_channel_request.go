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

// Request Object
type BatchInviteMembersToChannelRequest struct {
	Body *BatchInviteMembersToChannelRequestBody `json:"body,omitempty"`
}

func (o BatchInviteMembersToChannelRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchInviteMembersToChannelRequest", string(data)}, " ")
}
