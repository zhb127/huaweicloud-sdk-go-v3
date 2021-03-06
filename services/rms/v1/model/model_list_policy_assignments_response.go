/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPolicyAssignmentsResponse struct {
	// 规则列表
	Value          *[]PolicyAssignment `json:"value,omitempty"`
	PageInfo       *PageInfo           `json:"page_info,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListPolicyAssignmentsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPolicyAssignmentsResponse", string(data)}, " ")
}
