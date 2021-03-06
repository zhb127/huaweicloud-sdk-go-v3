/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateOrDeleteTagsRequest struct {
	InstanceId string                      `json:"instance_id"`
	Body       *CreateOrDeleteInstanceTags `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateOrDeleteTagsRequest", string(data)}, " ")
}
