/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 镜像导出请求体
type ExportImageRequestBody struct {
	// 目的文件的URL，格式：<bucket>:<file>。 说明：此处的OBS桶和镜像文件的存储类别必须是OBS标准存储。
	BucketUrl string `json:"bucket_url"`
	// 文件格式，支持qcow2、vhd、zvhd和vmdk。
	FileFormat ExportImageRequestBodyFileFormat `json:"file_format"`
}

func (o ExportImageRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ExportImageRequestBody", string(data)}, " ")
}

type ExportImageRequestBodyFileFormat struct {
	value string
}

type ExportImageRequestBodyFileFormatEnum struct {
	QCOW2 ExportImageRequestBodyFileFormat
	VHD   ExportImageRequestBodyFileFormat
	ZVHD  ExportImageRequestBodyFileFormat
	VMDK  ExportImageRequestBodyFileFormat
}

func GetExportImageRequestBodyFileFormatEnum() ExportImageRequestBodyFileFormatEnum {
	return ExportImageRequestBodyFileFormatEnum{
		QCOW2: ExportImageRequestBodyFileFormat{
			value: "qcow2",
		},
		VHD: ExportImageRequestBodyFileFormat{
			value: "vhd",
		},
		ZVHD: ExportImageRequestBodyFileFormat{
			value: "zvhd",
		},
		VMDK: ExportImageRequestBodyFileFormat{
			value: "vmdk",
		},
	}
}

func (c ExportImageRequestBodyFileFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ExportImageRequestBodyFileFormat) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}