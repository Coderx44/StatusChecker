package server

import "github.com/Coderx44/StatusChecker/service"

type dependencies struct {
	httpchecker service.StatusChecker
}

func InitDependencies() (dependencies, error) {
	statuschecker := service.NewHttpChecker()

	return dependencies{
		httpchecker: statuschecker,
	}, nil
}
