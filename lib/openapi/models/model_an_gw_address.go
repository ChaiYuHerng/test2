/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// describes the address of the access network gateway control node
type AnGwAddress struct {
	AnGwIpv4Addr string `json:"anGwIpv4Addr,omitempty" yaml:"anGwIpv4Addr" bson:"anGwIpv4Addr" mapstructure:"AnGwIpv4Addr"`
	AnGwIpv6Addr string `json:"anGwIpv6Addr,omitempty" yaml:"anGwIpv6Addr" bson:"anGwIpv6Addr" mapstructure:"AnGwIpv6Addr"`
}
