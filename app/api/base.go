package api

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code" `
	Msg  string      `json:"msg" `
	Data interface{} `json:"data" `
}

var SuccessResponse = &Response{
	Code: 200,
	Msg:  "success",
	Data: nil,
}
var ErrorRespnose = &Response{
	Code: 100,
	Msg:  "error",
	Data: nil,
}

type ResponseFn func(response *Response) *Response

// Code /*
func Code(code int) ResponseFn {
	return func(response *Response) *Response {
		response.Code = code
		return response
	}
}

// Err /*
func Err(err error) ResponseFn {
	return func(response *Response) *Response {
		if err != nil {
			response.Msg = err.Error()
		}
		return response
	}
}

// Msg /*
func Msg(msg string) ResponseFn {
	return func(response *Response) *Response {
		response.Msg = msg
		return response
	}
}

// Data /*
func Data(data interface{}) ResponseFn {
	return func(response *Response) *Response {
		response.Data = data
		return response
	}
}

// Success /*
func Success(ctx *gin.Context, fns ...ResponseFn) {
	var success = &Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	}
	if len(fns) > 0 {
		for _, fn := range fns {
			success = fn(success)
		}
	}
	ctx.JSON(200, success)
}

// Error /*
func Error(ctx *gin.Context, fns ...ResponseFn) {
	var err = &Response{
		Code: 100,
		Msg:  "error",
		Data: nil,
	}
	if len(fns) > 0 {
		for _, fn := range fns {
			err = fn(err)
		}
	}
	ctx.JSON(200, err)
}

// WeChatPayError /*
func WeChatPayError(ctx *gin.Context) {
	ctx.JSON(500, gin.H{
		"code":    "FAILED",
		"message": "支付失败",
	})
}
func WeChatPaySuccess(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    "SUCCESS",
		"message": "成功",
	})
}