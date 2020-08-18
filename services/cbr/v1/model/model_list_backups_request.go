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

// Request Object
type ListBackupsRequest struct {
	CheckpointId        *string                        `json:"checkpoint_id,omitempty"`
	Dec                 *bool                          `json:"dec,omitempty"`
	EndTime             *string                        `json:"end_time,omitempty"`
	ImageType           ListBackupsRequestImageType    `json:"image_type,omitempty"`
	Limit               *int32                         `json:"limit,omitempty"`
	Marker              *string                        `json:"marker,omitempty"`
	Name                *string                        `json:"name,omitempty"`
	Offset              *int32                         `json:"offset,omitempty"`
	ResourceAz          *string                        `json:"resource_az,omitempty"`
	ResourceId          *string                        `json:"resource_id,omitempty"`
	ResourceName        *string                        `json:"resource_name,omitempty"`
	ResourceType        ListBackupsRequestResourceType `json:"resource_type,omitempty"`
	Sort                *string                        `json:"sort,omitempty"`
	StartTime           *string                        `json:"start_time,omitempty"`
	Status              ListBackupsRequestStatus       `json:"status,omitempty"`
	VaultId             *string                        `json:"vault_id,omitempty"`
	EnterpriseProjectId *string                        `json:"enterprise_project_id,omitempty"`
	OwnType             ListBackupsRequestOwnType      `json:"own_type,omitempty"`
	MemberStatus        ListBackupsRequestMemberStatus `json:"member_status,omitempty"`
	ParentId            *string                        `json:"parent_id,omitempty"`
	UsedPercent         *string                        `json:"used_percent,omitempty"`
}

func (o ListBackupsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBackupsRequest", string(data)}, " ")
}

type ListBackupsRequestImageType struct {
	value string
}

type ListBackupsRequestImageTypeEnum struct {
	BACKUP      ListBackupsRequestImageType
	REPLICATION ListBackupsRequestImageType
}

func GetListBackupsRequestImageTypeEnum() ListBackupsRequestImageTypeEnum {
	return ListBackupsRequestImageTypeEnum{
		BACKUP: ListBackupsRequestImageType{
			value: "backup",
		},
		REPLICATION: ListBackupsRequestImageType{
			value: "replication",
		},
	}
}

func (c ListBackupsRequestImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListBackupsRequestImageType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type ListBackupsRequestResourceType struct {
	value string
}

type ListBackupsRequestResourceTypeEnum struct {
	OSCINDERVOLUME ListBackupsRequestResourceType
	OSNOVASERVER   ListBackupsRequestResourceType
}

func GetListBackupsRequestResourceTypeEnum() ListBackupsRequestResourceTypeEnum {
	return ListBackupsRequestResourceTypeEnum{
		OSCINDERVOLUME: ListBackupsRequestResourceType{
			value: "OS::Cinder::Volume",
		},
		OSNOVASERVER: ListBackupsRequestResourceType{
			value: "OS::Nova::Server",
		},
	}
}

func (c ListBackupsRequestResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListBackupsRequestResourceType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type ListBackupsRequestStatus struct {
	value string
}

type ListBackupsRequestStatusEnum struct {
	AVAILABLE       ListBackupsRequestStatus
	PROTECTING      ListBackupsRequestStatus
	DELETING        ListBackupsRequestStatus
	RESTORING       ListBackupsRequestStatus
	ERROR           ListBackupsRequestStatus
	WAITING_PROTECT ListBackupsRequestStatus
	WAITING_DELETE  ListBackupsRequestStatus
	WAITING_RESTORE ListBackupsRequestStatus
}

func GetListBackupsRequestStatusEnum() ListBackupsRequestStatusEnum {
	return ListBackupsRequestStatusEnum{
		AVAILABLE: ListBackupsRequestStatus{
			value: "available",
		},
		PROTECTING: ListBackupsRequestStatus{
			value: "protecting",
		},
		DELETING: ListBackupsRequestStatus{
			value: "deleting",
		},
		RESTORING: ListBackupsRequestStatus{
			value: "restoring",
		},
		ERROR: ListBackupsRequestStatus{
			value: "error",
		},
		WAITING_PROTECT: ListBackupsRequestStatus{
			value: "waiting_protect",
		},
		WAITING_DELETE: ListBackupsRequestStatus{
			value: "waiting_delete",
		},
		WAITING_RESTORE: ListBackupsRequestStatus{
			value: "waiting_restore",
		},
	}
}

func (c ListBackupsRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListBackupsRequestStatus) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type ListBackupsRequestOwnType struct {
	value string
}

type ListBackupsRequestOwnTypeEnum struct {
	ALL_GRANTED ListBackupsRequestOwnType
	PRIVATE     ListBackupsRequestOwnType
	SHARED      ListBackupsRequestOwnType
}

func GetListBackupsRequestOwnTypeEnum() ListBackupsRequestOwnTypeEnum {
	return ListBackupsRequestOwnTypeEnum{
		ALL_GRANTED: ListBackupsRequestOwnType{
			value: "all_granted",
		},
		PRIVATE: ListBackupsRequestOwnType{
			value: "private",
		},
		SHARED: ListBackupsRequestOwnType{
			value: "shared",
		},
	}
}

func (c ListBackupsRequestOwnType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListBackupsRequestOwnType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type ListBackupsRequestMemberStatus struct {
	value string
}

type ListBackupsRequestMemberStatusEnum struct {
	PENDING ListBackupsRequestMemberStatus
	ACCEPT  ListBackupsRequestMemberStatus
	REJECT  ListBackupsRequestMemberStatus
}

func GetListBackupsRequestMemberStatusEnum() ListBackupsRequestMemberStatusEnum {
	return ListBackupsRequestMemberStatusEnum{
		PENDING: ListBackupsRequestMemberStatus{
			value: "pending",
		},
		ACCEPT: ListBackupsRequestMemberStatus{
			value: "accept",
		},
		REJECT: ListBackupsRequestMemberStatus{
			value: "reject",
		},
	}
}

func (c ListBackupsRequestMemberStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListBackupsRequestMemberStatus) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}