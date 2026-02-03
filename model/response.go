package model

const SuccessReturnCode = "SUC0000"

type ResponseEntity[T any] struct {
	ReturnCode string `json:"returnCode"`
	ErrorMsg   string `json:"errorMsg"`
	Body       T      `json:"body"`
}

func Success() *ResponseEntity[any] {
	return &ResponseEntity[any]{
		ReturnCode: SuccessReturnCode,
		ErrorMsg:   "",
		Body:       nil,
	}
}

func SuccessWithBody[T any](body T) *ResponseEntity[T] {
	return &ResponseEntity[T]{
		ReturnCode: SuccessReturnCode,
		ErrorMsg:   "",
		Body:       body,
	}
}

func Failed(errorCode ErrorCode, args ...any) *ResponseEntity[any] {
	return &ResponseEntity[any]{
		ReturnCode: errorCode.Code(),
		ErrorMsg:   errorCode.Message(args...),
		Body:       nil,
	}
}

func FailedWithBody[T any](errorCode ErrorCode, body T, args ...any) *ResponseEntity[T] {
	return &ResponseEntity[T]{
		ReturnCode: errorCode.Code(),
		ErrorMsg:   errorCode.Message(args...),
		Body:       body,
	}
}

func IsSuccess(response *ResponseEntity[any]) bool {
	if response == nil {
		return false
	}
	return response.ReturnCode == SuccessReturnCode
}
