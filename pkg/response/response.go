package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct{
	Code int `json:"code"` //stt code
	Message string `json:"messsage"` //thong báo lỗi
	Data interface{} `json: "data"` //dữ liệu return
}

//success response
func SuccessRes( c*gin.Context, code int, data interface{}){
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Message: msg[code],
		Data: data,
	})
}

//error response

func ErrorRes ( c*gin.Context, code int, message string){
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Message: msg[code],
		Data: nil,
	})
}