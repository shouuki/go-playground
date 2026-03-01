package utility

import (
	"encoding/json"
	"fmt"
	"go-playground/model"
	"net/http"
)

// SendResponseJson marshals the given value v to JSON and writes it to the [http.ResponseWriter].
// It sets the Content-Type header to application/json
// If marshaling or writing the JSON fails, it returns an error.
func SendResponseJson(resp http.ResponseWriter, v any) error {
	resp.Header().Add("Content-Type", "application/json")

	bytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("marshal response json: %w", err)
	}
	if _, err := resp.Write(bytes); err != nil {
		return fmt.Errorf("write response body: %w", err)
	}
	return nil
}

// WriteResponseEntity wraps given body to a [model.ResponseEntity] with successful returnCode and then writes it to the [http.ResponseWriter].
func WriteResponseEntity(resp http.ResponseWriter, body ...any) error {
	if len(body) == 0 {
		return SendResponseJson(resp, model.Success())
	}
	return SendResponseJson(resp, model.SuccessWithBody(body[0]))
}

// WriteFailedResponseEntity wraps given body to a [model.ResponseEntity] with failed returnCode and then writes it to the [http.ResponseWriter].
//
// Field named errorMsg in [model.ResponseEntity] is calculated from the given [error].
func WriteFailedResponseEntity(resp http.ResponseWriter, err error, body ...any) error {
	var entity *model.ResponseEntity[any]
	if appError, ok := err.(*model.AppError); ok {
		entity = model.Failed(appError.ErrorCode(), appError.Args()...)
	} else {
		entity = model.Failed(model.SystemErrorCode, err.Error())
	}
	if len(body) > 0 {
		entity.Body = body[0]
	}
	return SendResponseJson(resp, entity)
}
