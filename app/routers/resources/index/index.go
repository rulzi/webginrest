package index

import (
	controller "webginrest/app/pkg/controller/index"
)

// Controller set index controller
func Controller() controller.Index {
	controller := controller.IndexController()

	return controller
}
