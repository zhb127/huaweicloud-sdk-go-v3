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
type CreateOrUpdateTagsResponse struct {
}

func (o CreateOrUpdateTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateOrUpdateTagsResponse", string(data)}, " ")
}