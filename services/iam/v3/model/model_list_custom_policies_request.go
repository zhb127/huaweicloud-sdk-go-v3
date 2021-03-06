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
type ListCustomPoliciesRequest struct {
	Page    *int32 `json:"page,omitempty"`
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListCustomPoliciesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListCustomPoliciesRequest", string(data)}, " ")
}
