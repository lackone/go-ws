package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Return struct {
	Code int         `json:"code"` //状态码
	Msg  string      `json:"msg"`  //消息
	Data interface{} `json:"data"` //数据
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToSuccess(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, Return{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}

func (r *Response) ToError(err *errcode.Error) {
	r.Ctx.JSON(err.StatusCode(), Return{
		Code: err.Code(),
		Msg:  err.Msg(),
		Data: err.Details(),
	})
}
