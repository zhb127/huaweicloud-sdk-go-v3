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

// 创建包周期的弹性公网IP
type CreatePrePaidPublicipRequestBody struct {
	Publicip    *CreatePrePaidPublicipOption            `json:"publicip"`
	Bandwidth   *CreatePublicipBandwidthOption          `json:"bandwidth"`
	ExtendParam *CreatePrePaidPublicipExtendParamOption `json:"extendParam,omitempty"`
}

func (o CreatePrePaidPublicipRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePrePaidPublicipRequestBody", string(data)}, " ")
}
