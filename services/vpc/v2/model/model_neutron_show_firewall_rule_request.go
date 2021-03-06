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
type NeutronShowFirewallRuleRequest struct {
	FirewallRuleId string `json:"firewall_rule_id"`
}

func (o NeutronShowFirewallRuleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronShowFirewallRuleRequest", string(data)}, " ")
}
