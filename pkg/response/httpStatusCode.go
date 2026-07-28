package response

const (
	ErrCodeSucccess = 20001 //Success
	ErrCodeParamsInvalid = 20003 // emai is invalid

)

//message 
var msg = map[int] string{
	ErrCodeSucccess:"success",
	ErrCodeParamsInvalid:"email is invalid",
}

