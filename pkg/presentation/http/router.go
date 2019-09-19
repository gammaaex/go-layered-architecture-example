package http

import (
	"github.com/go-chi/chi"
	"go-layered-architecture-example/pkg/presentation/http/handler"
	"go-layered-architecture-example/pkg/registry"
	"net/http"
)

func Register(r *chi.Mux) {
	repository := registry.NewRepository()
	service := registry.NewService()
	usecase := registry.NewUsecase()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This Go app work on CGI"))
	})

	userRepository := repository.NewUser()
	userService := service.NewUser(userRepository)
	userUsecase := usecase.NewUser(userService)
	userHandler := handler.NewUser(userUsecase)
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
	})
}
