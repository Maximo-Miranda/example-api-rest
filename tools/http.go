package tools

import (
	"net/http"
	"encoding/json"
	"strings"
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

// CheckedIfKeyExist ...
func CheckedIfKeyExist(in interface{}, fieldName string) (bool, error) {

	var mapper map[string]interface{}

	inSerialized, err := json.Marshal(in)
	if err != nil {
		return false, nil
	}

	err = json.Unmarshal(inSerialized, &mapper)
	if err != nil {
		return false, err
	}

	for key := range mapper {
		if strings.ToUpper(key) == strings.ToUpper(fieldName) {
			return true, nil
		}
	}

	return false, nil
}