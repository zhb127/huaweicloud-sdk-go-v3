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
type ShowL7ruleRequest struct {
	L7policyId string `json:"l7policy_id"`
	L7ruleId   string `json:"l7rule_id"`
}

func (o ShowL7ruleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowL7ruleRequest", string(data)}, " ")
}
