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

// Request Object
type AssociateAgencyWithProjectPermissionRequest struct {
	ProjectId string `json:"project_id"`
	AgencyId  string `json:"agency_id"`
	RoleId    string `json:"role_id"`
}

func (o AssociateAgencyWithProjectPermissionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AssociateAgencyWithProjectPermissionRequest", string(data)}, " ")
}