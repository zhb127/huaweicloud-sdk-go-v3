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

// Request Object
type ShowPoolRequest struct {
	PoolId string `json:"pool_id"`
}

func (o ShowPoolRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPoolRequest", string(data)}, " ")
}
