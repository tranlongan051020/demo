package user

import (
	userUc "demo/internal/usecase/user"
	"net/http"
)

type UserCtrl interface {
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type instance struct {
	userUc userUc.UserUsecase
}

func New(userUc userUc.UserUsecase) UserCtrl {
	return &instance{
		userUc: userUc,
	}
}
