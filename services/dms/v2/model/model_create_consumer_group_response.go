/*
 * DMS
 *
 * DMS Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateConsumerGroupResponse struct {
	// 消费组信息。
	Groups         *[]CreateConsumerGroupRespGroups `json:"groups,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o CreateConsumerGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateConsumerGroupResponse", string(data)}, " ")
}
