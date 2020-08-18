/*
 * cloudide
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

type InstanceEdgeParam struct {
	// cpu架构 x86|arm
	Arch InstanceEdgeParamArch `json:"arch,omitempty"`
	// cpu规格.arm架构支持4U8G，x86架构支持1U1G,2U4G,2U8G 与技术栈配置的规格对应，可通过技术栈管理ListStacksByTag接口获取。如果标签不为空，以标签配置的技术栈规格为准。 quantum技术栈，x86架构cpu规格为2U8G;其他技术栈，x86架构cpu规格为1U1G,2U4G
	CpuMemory InstanceEdgeParamCpuMemory `json:"cpu_memory"`
	// 描述。长度不操过100个字符
	Description *string `json:"description,omitempty"`
	// 实例名。 可以输入中文、数字、字母、下划线、点、破折号。长度介于3-100之间
	InstanceName string `json:"instance_name"`
	// 组织id（对应华为云账号的domainId）
	InstanceUserDomainId *string `json:"instance_user_domain_id,omitempty"`
	// 用户组织名
	InstanceUserDomainName string `json:"instance_user_domain_name"`
	// 用户id
	InstanceUserId *string `json:"instance_user_id,omitempty"`
	// 用户名
	InstanceUserName string `json:"instance_user_name"`
	// 是否临时实例。false页面会显示
	IsTemporary *bool `json:"is_temporary,omitempty"`
	// 插件列表
	Plugins []Plugin `json:"plugins,omitempty"`
	// PVC规格 5GB|10GB|20GB
	PvcQuantity InstanceEdgeParamPvcQuantity `json:"pvc_quantity"`
	// 实例的生命周期 arm架构,生命周期只能设置成30，60。x86架构可取值为30，60，240，1440和-1。除-1外，其它值的单位为“分钟”。实例在到达生命周期后，将会被暂停（已保存的数据不会被删除）。-1表示实例不会自动停止。
	RefreshTime *string `json:"refresh_time,omitempty"`
	// 技术栈ID 目前可取值all，java，go，python，cpp，nodejs，quantum，blockchain，dcn，vue，ruby。
	StackId string `json:"stack_id"`
}

func (o InstanceEdgeParam) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"InstanceEdgeParam", string(data)}, " ")
}

type InstanceEdgeParamArch struct {
	value string
}

type InstanceEdgeParamArchEnum struct {
	X86 InstanceEdgeParamArch
	ARM InstanceEdgeParamArch
}

func GetInstanceEdgeParamArchEnum() InstanceEdgeParamArchEnum {
	return InstanceEdgeParamArchEnum{
		X86: InstanceEdgeParamArch{
			value: "x86",
		},
		ARM: InstanceEdgeParamArch{
			value: "arm",
		},
	}
}

func (c InstanceEdgeParamArch) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstanceEdgeParamArch) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type InstanceEdgeParamCpuMemory struct {
	value string
}

type InstanceEdgeParamCpuMemoryEnum struct {
	_1_U1_G InstanceEdgeParamCpuMemory
	_2_U4_G InstanceEdgeParamCpuMemory
	_2_U8_G InstanceEdgeParamCpuMemory
	_4_U8_G InstanceEdgeParamCpuMemory
}

func GetInstanceEdgeParamCpuMemoryEnum() InstanceEdgeParamCpuMemoryEnum {
	return InstanceEdgeParamCpuMemoryEnum{
		_1_U1_G: InstanceEdgeParamCpuMemory{
			value: "1U1G",
		},
		_2_U4_G: InstanceEdgeParamCpuMemory{
			value: "2U4G",
		},
		_2_U8_G: InstanceEdgeParamCpuMemory{
			value: "2U8G",
		},
		_4_U8_G: InstanceEdgeParamCpuMemory{
			value: "4U8G",
		},
	}
}

func (c InstanceEdgeParamCpuMemory) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstanceEdgeParamCpuMemory) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type InstanceEdgeParamPvcQuantity struct {
	value string
}

type InstanceEdgeParamPvcQuantityEnum struct {
	_5_GB  InstanceEdgeParamPvcQuantity
	_10_GB InstanceEdgeParamPvcQuantity
	_20_GB InstanceEdgeParamPvcQuantity
}

func GetInstanceEdgeParamPvcQuantityEnum() InstanceEdgeParamPvcQuantityEnum {
	return InstanceEdgeParamPvcQuantityEnum{
		_5_GB: InstanceEdgeParamPvcQuantity{
			value: "5GB",
		},
		_10_GB: InstanceEdgeParamPvcQuantity{
			value: "10GB",
		},
		_20_GB: InstanceEdgeParamPvcQuantity{
			value: "20GB",
		},
	}
}

func (c InstanceEdgeParamPvcQuantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstanceEdgeParamPvcQuantity) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}