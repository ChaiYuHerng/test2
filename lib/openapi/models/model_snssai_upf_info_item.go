/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SnssaiUpfInfoItem struct {
	SNssai         *Snssai          `json:"sNssai" yaml:"sNssai" bson:"sNssai" mapstructure:"SNssai"`
	DnnUpfInfoList []DnnUpfInfoItem `json:"dnnUpfInfoList" yaml:"dnnUpfInfoList" bson:"dnnUpfInfoList" mapstructure:"DnnUpfInfoList"`
}
