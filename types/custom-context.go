package types

import (
	"fmt"
	"net/http"

	"github.com/antigloss/go/logger"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Custom request context implementation of gin context
type TemplateRequestContext struct {
	Context   *gin.Context
	RequestId string
	//User      string
}

// Convert incoming gin request to custom context
func ConvertContext(c *gin.Context) *TemplateRequestContext {
	t := new(TemplateRequestContext)
	t.Context = c
	id, _ := uuid.NewV4()
	t.RequestId = id.String()
	return t
}

func (t *TemplateRequestContext) PrintString() string {
	return fmt.Sprintf("%s :  ")
}

// All error response should go via this call
func (t *TemplateRequestContext) ERROR(err error, errorCode int, abortRequest bool) {
	if err != nil {
		if !t.Context.IsAborted() {
			if err != nil {
				errStr := fmt.Sprintf("Error occured while handling request for %s:%s path received error %s", t.Context.Request.Method, t.Context.Request.URL.Path, err.Error())
				logger.Error(errStr)
				if abortRequest {
					resp := gin.H{
						"status":         errorCode,
						"data":           "",
						"message":        err.Error(),
						"request_path":   t.Context.Request.URL.Path,
						"request_method": t.Context.Request.Method,
					}
					t.Context.AbortWithStatusJSON(errorCode, resp)
				}
			}
		}
	}
}

func (t *TemplateRequestContext) CUSTOMERROR(msg string, errorCode int) {

	errStr := fmt.Sprintf("Error occured while handling request for %s:%s path received error %s", t.Context.Request.Method, t.Context.Request.URL.Path, msg)
	logger.Error(errStr)
	resp := gin.H{
		"status":         errorCode,
		"data":           "",
		"message":        msg,
		"request_path":   t.Context.Request.URL.Path,
		"request_method": t.Context.Request.Method,
	}
	t.Context.AbortWithStatusJSON(errorCode, resp)
}

// Request will be aborted by default
func (t *TemplateRequestContext) FATAL(err error, errorCode int) {
	t.ERROR(err, errorCode, true)
}

// All positive response should go via this call
func (t *TemplateRequestContext) OK(data interface{}, message interface{}) {
	resp := gin.H{"status": http.StatusOK, "data": data, "message": message}
	if !t.Context.IsAborted() {
		t.Context.JSON(http.StatusOK, resp)
	}
}

// Custom Info for logger
func (t *TemplateRequestContext) INFO(msg string) {
	logger.Info("%s : %s - %s", t.Context.Request.Method, t.Context.Request.URL.Path, msg)
}

func (t *TemplateRequestContext) GetUrlParam(param string) string {
	return t.Context.Param(param)
}

func (t *TemplateRequestContext) GetUrlQueryParam(key string) string {
	return t.Context.Query(key)
}
