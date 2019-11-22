package tools

import (
	"net/http"
	"encoding/json"
)

// ResponseHandlers ...
func ResponseHandlers(res http.ResponseWriter, data interface{}, err interface{}, statusCode int32) {

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(int(statusCode))

	str := struct {
		StatusCode int32       `json:"status_code"`
		Response   interface{} `json:"response"`
		Error      interface{} `json:"error"`
	}{
		StatusCode: statusCode,
		Response:   data,
		Error:      err,
	}

	serialized, err := json.Marshal(str)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Serialization Error"))
		return
	}

	res.Write(serialized)
}