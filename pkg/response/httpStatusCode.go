package response

const (
	ErrCodeSucccess = 20001 //Success
	ErrCodeParamsInvalid = 20003 // emai is invalid
	ErrInvalid = 30003 // email in valid
)

//message 
var msg = map[int] string{
	ErrCodeSucccess:"success",
	ErrCodeParamsInvalid:"email is invalid",
	ErrInvalid:"token is invalid",

}


