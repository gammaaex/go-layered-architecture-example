package handler

import (
	"encoding/json"
	"go-layered-architecture-example/pkg/application/input"
	"go-layered-architecture-example/pkg/domain/service"
	"go-layered-architecture-example/pkg/presentation/http/request"
	"go-layered-architecture-example/pkg/presentation/http/response"
	"net/http"
)

type User struct {
	service service.User
}

func NewUser(service service.User) *User {
	return &User{service: service}
}

func (h *User) Create(w http.ResponseWriter, r *http.Request) {
	var req request.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	in := &input.UserCreate{}
	req.Write(in)

	out, err := h.service.Create(in)
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(response.NewUser(out))
}
