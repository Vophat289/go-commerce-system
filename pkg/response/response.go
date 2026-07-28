package respone

type ResponseData struct{
	Code int `json:"code"` //stt code
	Message string `json:"messsage"` //thong báo lỗi
	Data interface{} `json: "data"` //dữ liệu return
}