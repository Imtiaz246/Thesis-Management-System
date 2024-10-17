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

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	regCodes, found := codeReg[ErrSuccess]
	if !found {
		resp := Response{Code: 0, Message: "", Data: data}
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp := Response{Code: regCodes.serverCode, Message: ErrSuccess.Error(), Data: data}
	ctx.JSON(regCodes.httpCode, resp)
}

func HandleError(ctx *gin.Context, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}

	regCodes, found := codeReg[err]
	if !found {
		resp := Response{Code: 0, Message: "Unknown Error", Data: data}
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := Response{Code: regCodes.serverCode, Message: err.Error(), Data: data}
	ctx.JSON(regCodes.httpCode, resp)
}

type Error struct {
	Code    int
	Message string
}

type errCodes struct {
	httpCode   int
	serverCode int
}

var codeReg = map[error]errCodes{}

func registerCodes(serverCode, httpCode int, msg string) error {
	err := errors.New(msg)
	codeReg[err] = errCodes{
		serverCode: serverCode,
		httpCode:   httpCode,
	}
	return err
}

func (e Error) Error() string {
	return e.Message
}
