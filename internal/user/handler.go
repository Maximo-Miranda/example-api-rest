package user

import (
	"net/http"
	"encoding/json"
	"github.com/Maximo-Miranda/example-api-rest/tools"
)

// StoreUser ...
func StoreUser(w http.ResponseWriter, r *http.Request) {

	model := User{}

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		tools.ResponseHandlers(w, nil, "Error json decode: the body request is null or invalid datatype", http.StatusUnprocessableEntity)
		return
	}

	err = validateCreateClientRequest(&model)
	if err != nil {
		tools.ResponseHandlers(w, nil, err, http.StatusBadRequest)
		return
	}

	userDB, err := model.Save()
	if err != nil {
		tools.ResponseHandlers(w, nil, err, http.StatusInternalServerError)
		return
	}

	tools.ResponseHandlers(w, userDB, nil, http.StatusOK)
}
