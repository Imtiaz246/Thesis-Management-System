package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}, msgs ...successMsg) {
	if data == nil {
		data = map[string]interface{}{}
	}
	var msg successMsg
	if len(msg) > 0 {
		msg = msgs[0]
	}
	regCodes, found := successCodeRegistry[msg]
	if !found {
		resp := Response{Code: http.StatusOK, Message: "", Data: data}
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp := Response{Code: regCodes.serverCode, Message: msg.String(), Data: data}
	ctx.JSON(regCodes.httpCode, resp)
}

func HandleError(ctx *gin.Context, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	regCodes, found := errCodeRegistry[err]
	if !found {
		var svrErr serverError
		if errors.As(err, &svrErr) {
			resp := Response{Code: svrErr.code, Message: svrErr.message, Data: data}
			ctx.JSON(svrErr.code, resp)
			return
		}

		resp := Response{Code: 0, Message: err.Error(), Data: data}
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := Response{Code: regCodes.serverCode, Message: err.Error(), Data: data}
	ctx.JSON(regCodes.httpCode, resp)
}

type serverError struct {
	code    int
	message string
}

func (e serverError) Error() string {
	return e.message
}

func ServerError(httpCode int, err error) error {
	return serverError{
		code:    httpCode,
		message: err.Error(),
	}
}

type successMsg string

func (s successMsg) String() string {
	return string(s)
}

type codes struct {
	httpCode   int
	serverCode int
}

var (
	successCodeRegistry = map[successMsg]codes{}
	errCodeRegistry     = map[error]codes{}
)

func regSuccessCode(serverCode, httpCode int, msg string) successMsg {
	success := successMsg(msg)
	successCodeRegistry[success] = codes{
		serverCode: serverCode,
		httpCode:   httpCode,
	}
	return successMsg(msg)
}

func regErrCode(serverCode, httpCode int, msg string) error {
	err := errors.New(msg)
	errCodeRegistry[err] = codes{
		serverCode: serverCode,
		httpCode:   httpCode,
	}
	return err
}
