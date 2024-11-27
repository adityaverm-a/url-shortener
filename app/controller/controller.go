package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	errorCodeKey    = "error_code"
	codeKey         = "code"
	codeSuccess     = "success"
	codeFailed      = "failed"
	dataKey         = "data"
	msgKey          = "msg"
	errorMessageKey = "errorMessage"
)

type Controller struct {
}

func (ctrl Controller) Send(gCtx *gin.Context, data interface{}) {
	ctrl.sendResponse(http.StatusOK, gCtx, gin.H{codeKey: codeSuccess, dataKey: data})
}

func (ctrl Controller) SendWithError(gCtx *gin.Context, err error) {
	// httpStatusCode := ctrl.getHTTPStatusCodeFromError(err)
	httpStatusCode := 200

	// passing context of request because that's where request id is stored.
	ctrl.logError(err)

	ctrl.sendResponse(httpStatusCode, gCtx, gin.H{codeKey: codeFailed, msgKey: err.Error()})
}

func (ctrl Controller) sendResponse(statusCode int, gCtx *gin.Context, ginH gin.H) {
	gCtx.JSON(statusCode, ginH)
}

func (ctrl Controller) logError(err error) {
	log.Println(err.Error())
}
