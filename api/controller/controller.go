package controller

import (
	userCtrl "demo/api/controller/user"
)

type Controller struct {
	UserCtrl userCtrl.UserCtrl
}

func New(
	UserCtrl userCtrl.UserCtrl,
) *Controller {
	return &Controller{
		UserCtrl: UserCtrl,
	}
}
