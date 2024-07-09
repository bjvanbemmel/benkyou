package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response[T any] struct {
	Data T `json:"data"`
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func New[T any](w http.ResponseWriter, status int, object T) {
	raw, _ := json.Marshal(Response[T]{
		Data: object,
	})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	fmt.Fprint(w, string(raw))
}

func NewError(w http.ResponseWriter, status int, err error) {
	raw, _ := json.Marshal(Error{
		Message: err.Error(),
		Status:  status,
	})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprint(w, string(raw))
}
