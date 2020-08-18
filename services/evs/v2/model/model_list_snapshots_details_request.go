/*
 * EVS
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
type ListSnapshotsDetailsRequest struct {
	Offset               *int32  `json:"offset,omitempty"`
	Limit                *int32  `json:"limit,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Status               *string `json:"status,omitempty"`
	VolumeId             *string `json:"volume_id,omitempty"`
	AvailabilityZone     *string `json:"availability_zone,omitempty"`
	Id                   *string `json:"id,omitempty"`
	DedicatedStorageName *string `json:"dedicated_storage_name,omitempty"`
	DedicatedStorageId   *string `json:"dedicated_storage_id,omitempty"`
	ServiceType          *string `json:"service_type,omitempty"`
	EnterpriseProjectId  *string `json:"enterprise_project_id,omitempty"`
}

func (o ListSnapshotsDetailsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSnapshotsDetailsRequest", string(data)}, " ")
}