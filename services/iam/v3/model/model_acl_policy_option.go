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

//
type AclPolicyOption struct {
	// 允许访问的IP地址或网段。
	AllowAddressNetmasks []AllowAddressNetmasksOption `json:"allow_address_netmasks"`
	// 允许访问的IP地址区间
	AllowIpRanges []AllowIpRangesOption `json:"allow_ip_ranges"`
}

func (o AclPolicyOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AclPolicyOption", string(data)}, " ")
}
