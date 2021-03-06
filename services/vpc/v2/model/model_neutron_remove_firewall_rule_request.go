/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronRemoveFirewallRuleRequest struct {
	FirewallPolicyId string                                `json:"firewall_policy_id"`
	Body             *NeutronRemoveFirewallRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronRemoveFirewallRuleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronRemoveFirewallRuleRequest", string(data)}, " ")
}
