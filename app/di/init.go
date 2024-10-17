package di

import (
	"ginTemplate/internal/controller"
)

type Initialization struct {
	TestController controller.TestController
}

func NewInitialization(testController controller.TestController) *Initialization {
	init := &Initialization{
		TestController: testController,
	}

	return init
}
