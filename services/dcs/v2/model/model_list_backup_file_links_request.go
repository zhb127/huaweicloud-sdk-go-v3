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

// Request Object
type ListBackupFileLinksRequest struct {
	InstanceId string                  `json:"instance_id"`
	BackupId   string                  `json:"backup_id"`
	Body       *DownloadBackupFilesReq `json:"body,omitempty"`
}

func (o ListBackupFileLinksRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBackupFileLinksRequest", string(data)}, " ")
}
