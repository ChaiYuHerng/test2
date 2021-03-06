/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package uecontextmanagement

import (
	// "fmt"
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/udm/logger"
	"free5gc/src/udm/handler"
	udm_message "free5gc/src/udm/handler/message"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegistrationAmf3gppAccess - register as AMF for 3GPP access
func RegistrationAmf3gppAccess(c *gin.Context) {
	var amf3GppAccessRegistration models.Amf3GppAccessRegistration
	if err := c.ShouldBindJSON(&amf3GppAccessRegistration); err != nil {
		logger.UecmLog.Errorln(err)
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, amf3GppAccessRegistration)
	req.Params["ueId"] = c.Param("ueId")

	handlerMsg := udm_message.NewHandlerMessage(udm_message.EventRegistrationAmf3gppAccess, req)
	handler.SendMessage(handlerMsg)
	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse
	c.Header("Location", HTTPResponse.Header.Get("Location"))
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
	return
}
