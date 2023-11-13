package login

import (
	controller "webginrest/app/pkg/controller/login"
	repository "webginrest/app/pkg/repository/login"
	usecase "webginrest/app/pkg/usecase/login"
)

// Controller set login controller
func Controller() controller.Login {
	repository := repository.Repository()
	usecase := usecase.Usecase(repository)
	controller := controller.LoginController(usecase)

	return controller
}
