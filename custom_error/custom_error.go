package custom_error

// type CustomError struct {
// 	Code    int32  `json:"code"`
// 	Message string `json:"message"`
// }

// func (custom CustomError) Error() string {
// 	return custom.Message
// }

type SystemError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (system SystemError) Error() string {
	return system.Message
}

type BusinessError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (business BusinessError) Error() string {
	return business.Message
}
