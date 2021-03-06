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
type ListL7PoliciesRequest struct {
	Action              *[]string `json:"action,omitempty"`
	AdminStateUp        *bool     `json:"admin_state_up,omitempty"`
	Description         *[]string `json:"description,omitempty"`
	DisplayAllRules     *bool     `json:"display_all_rules,omitempty"`
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	Id                  *[]string `json:"id,omitempty"`
	Limit               *int32    `json:"limit,omitempty"`
	ListenerId          *[]string `json:"listener_id,omitempty"`
	Marker              *string   `json:"marker,omitempty"`
	Name                *[]string `json:"name,omitempty"`
	PageReverse         *bool     `json:"page_reverse,omitempty"`
	Position            *[]int32  `json:"position,omitempty"`
	ProvisioningStatus  *[]string `json:"provisioning_status,omitempty"`
	RedirectListenerId  *[]string `json:"redirect_listener_id,omitempty"`
	RedirectPoolId      *[]string `json:"redirect_pool_id,omitempty"`
	RedirectUrl         *[]string `json:"redirect_url,omitempty"`
}

func (o ListL7PoliciesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListL7PoliciesRequest", string(data)}, " ")
}
