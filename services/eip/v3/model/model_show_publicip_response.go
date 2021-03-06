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
type ShowPublicipResponse struct {
	// 本次请求的编号
	RequestId      *string           `json:"request_id,omitempty"`
	Publicip       *PublicipShowResp `json:"publicip,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowPublicipResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPublicipResponse", string(data)}, " ")
}
