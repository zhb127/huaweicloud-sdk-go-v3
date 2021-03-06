/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ChangeInstanceResponse struct {
	// Job ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ChangeInstanceResponse", string(data)}, " ")
}
