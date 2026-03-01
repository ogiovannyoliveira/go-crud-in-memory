package api

import (
	"encoding/json"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func Insert(app Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser User

		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			SendJSON(w, Response{Error: "Could not parse body."}, http.StatusUnprocessableEntity)
			return
		}

		if err := newUser.Validate(); err != nil {
			SendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		id := ID(uuid.NewV4())
		app.Data[id] = newUser

		SendJSON(
			w,
			Response{Data: NewUserResponse(id, newUser)},
			http.StatusCreated,
		)
	}
}

func FindAll(app Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := make([]UserResponse, 0, len(app.Data))

		for id, body := range app.Data {
			users = append(users, NewUserResponse(id, body))
		}

		SendJSON(w, Response{Data: users}, http.StatusOK)
	}
}
