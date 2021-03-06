/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type RestoreInstanceBody struct {
	// 备份记录ID。
	BackupId string `json:"backup_id"`
	// 恢复缓存实例的备注信息。
	Remark *string `json:"remark,omitempty"`
}

func (o RestoreInstanceBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestoreInstanceBody", string(data)}, " ")
}
