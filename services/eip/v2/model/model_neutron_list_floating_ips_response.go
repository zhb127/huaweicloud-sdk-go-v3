/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NeutronListFloatingIpsResponse struct {
	// floatingip对象列表
	Floatingips *[]FloatingIpResp `json:"floatingips,omitempty"`
	// marker分页结构
	FloatingipsLinks *[]Pager `json:"floatingips_links,omitempty"`
	HttpStatusCode   int      `json:"-"`
}

func (o NeutronListFloatingIpsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronListFloatingIpsResponse", string(data)}, " ")
}
