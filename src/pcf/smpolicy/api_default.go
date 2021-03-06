/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smpolicy

import (
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/pcf/handler/message"
	"free5gc/src/pcf/logger"
	"free5gc/src/pcf/util"

	"github.com/gin-gonic/gin"
)

// SmPoliciesPost -
func SmPoliciesPost(c *gin.Context) {

	var smPolicyContextData models.SmPolicyContextData
	err := c.ShouldBindJSON(&smPolicyContextData)
	if err != nil {
		rsp := util.GetProblemDetail("Malformed request syntax", util.ERROR_INITIAL_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}
	if smPolicyContextData.Supi == "" || smPolicyContextData.SliceInfo == nil || len(smPolicyContextData.SliceInfo.Sd) != 6 {
		rsp := util.GetProblemDetail("Errorneous/Missing Mandotory IE", util.ERROR_INITIAL_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, smPolicyContextData)
	channelMsg := message.NewHttpChannelMessage(message.EventSMPolicyCreate, req)

	message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse

	for key, val := range HTTPResponse.Header {
		c.Header(key, val[0])
	}
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// SmPoliciesSmPolicyIdDeletePost -
func SmPoliciesSmPolicyIdDeletePost(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["smPolicyId"] = c.Params.ByName("smPolicyId")
	channelMsg := message.NewHttpChannelMessage(message.EventSMPolicyDelete, req)
	message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// SmPoliciesSmPolicyIdGet -
func SmPoliciesSmPolicyIdGet(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["smPolicyId"] = c.Params.ByName("smPolicyId")
	channelMsg := message.NewHttpChannelMessage(message.EventSMPolicyGet, req)

	message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// SmPoliciesSmPolicyIdUpdatePost -
func SmPoliciesSmPolicyIdUpdatePost(c *gin.Context) {
	var smPolicyUpdateContextData models.SmPolicyUpdateContextData
	c.ShouldBindJSON(&smPolicyUpdateContextData)
	req := http_wrapper.NewRequest(c.Request, smPolicyUpdateContextData)
	req.Params["smPolicyId"] = c.Params.ByName("smPolicyId")
	channelMsg := message.NewHttpChannelMessage(message.EventSMPolicyUpdate, req)

	message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
