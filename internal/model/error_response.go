package model

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
	Err     error  `json:"-"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
