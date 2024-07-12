package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bjvanbemmel/benkyou/internal/errors"
)

var (
	target string = os.Getenv("API_BUILD_TARGET")
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
	if status == http.StatusInternalServerError && target == "release" {
		err = errors.ErrSomethingWentWrong
	}

	raw, _ := json.Marshal(Error{
		Message: err.Error(),
		Status:  status,
	})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, "%s", raw)
}
