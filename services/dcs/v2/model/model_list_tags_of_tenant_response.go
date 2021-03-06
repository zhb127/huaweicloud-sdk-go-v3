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

// Response Object
type ListTagsOfTenantResponse struct {
	// 标签列表。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTagsOfTenantResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTagsOfTenantResponse", string(data)}, " ")
}
