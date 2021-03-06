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
type ListPoolsRequest struct {
	AdminStateUp                   *bool     `json:"admin_state_up,omitempty"`
	Description                    *[]string `json:"description,omitempty"`
	EnterpriseProjectId            *[]string `json:"enterprise_project_id,omitempty"`
	HealthmonitorId                *[]string `json:"healthmonitor_id,omitempty"`
	Id                             *[]string `json:"id,omitempty"`
	IpVersion                      *[]string `json:"ip_version,omitempty"`
	LbAlgorithm                    *[]string `json:"lb_algorithm,omitempty"`
	Limit                          *int32    `json:"limit,omitempty"`
	LoadbalancerId                 *[]string `json:"loadbalancer_id,omitempty"`
	Marker                         *string   `json:"marker,omitempty"`
	MemberAddress                  *[]string `json:"member_address,omitempty"`
	MemberDeletionProtectionEnable *bool     `json:"member_deletion_protection_enable,omitempty"`
	MemberDeviceId                 *[]string `json:"member_device_id,omitempty"`
	Name                           *[]string `json:"name,omitempty"`
	PageReverse                    *bool     `json:"page_reverse,omitempty"`
	Protocol                       *[]string `json:"protocol,omitempty"`
}

func (o ListPoolsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPoolsRequest", string(data)}, " ")
}
