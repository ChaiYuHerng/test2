/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"free5gc/lib/http_wrapper"
	"free5gc/src/udr/handler/message"

	"github.com/gin-gonic/gin"
)

// QueryTraceData - Retrieves the trace configuration data of a UE
func QueryTraceData(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")
	req.Params["servingPlmnId"] = c.Params.ByName("servingPlmnId")

	handlerMsg := message.NewHandlerMessage(message.EventQueryTraceData, req)
	message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
