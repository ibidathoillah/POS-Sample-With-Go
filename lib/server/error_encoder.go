package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	errors "github.com/ibidathoillah/majoo-test/lib/errors"
)

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	code := http.StatusInternalServerError
	message := "Something Went Wrong"
	messageList := []string{}

	if sc, ok := err.(*errors.Error); ok {
		code = sc.StatusCode
		message = sc.Message
	}

	switch err.(type) {

	// When catch error validation struct
	case validator.ValidationErrors:

		w.WriteHeader(http.StatusBadRequest)
		errValidation := err.(validator.ValidationErrors)
		for _, e := range errValidation {
			messageList = append(messageList, fmt.Sprintf("%s is %s", e.StructField(), e.Tag()))
		}
		message = messageList[0]

		break

	default:
		w.WriteHeader(code)
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error":        message,
		"message":      err.Error(),
		"message_list": messageList,
	})
}
