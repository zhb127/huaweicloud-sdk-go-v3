/*
 * BCS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 邀请联盟成员
type BatchInviteMembersToChannelRequestBody struct {
	// 邀请实例id
	BcsId *string `json:"bcs_id,omitempty"`
	// 邀请加入的通道名
	ChannelName *string `json:"channel_name,omitempty"`
	// 发出邀请的租户名
	InvitorUsername *string `json:"invitor_username,omitempty"`
	// 被邀请的用户列表
	InvitedUserinfo *[]InvitedDomain `json:"invited_userinfo,omitempty"`
}

func (o BatchInviteMembersToChannelRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchInviteMembersToChannelRequestBody", string(data)}, " ")
}
