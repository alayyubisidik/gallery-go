package helper

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result any) error {
	decoder := json.NewDecoder(request.Body)
	// decoder.DisallowUnknownFields()

	// Check if the request body is empty
	if request.Body == nil || request.ContentLength == 0 {
		return errors.New("request body is empty")
	}

	err := decoder.Decode(result)
	if err != nil {
		if err == io.EOF {
			return errors.New("request body is empty")
		}
		return err
	}

	return nil
}

func WriteToResponseBody(writer http.ResponseWriter, response any, statusCode int) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}