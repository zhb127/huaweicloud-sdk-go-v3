/*
 * Cbr
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
type VaultMigrateResourceReq struct {
	// 目标存储库
	DestinationVaultId string `json:"destination_vault_id"`
	// 待迁移的资源ID
	ResourceIds []string `json:"resource_ids"`
}

func (o VaultMigrateResourceReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"VaultMigrateResourceReq", string(data)}, " ")
}