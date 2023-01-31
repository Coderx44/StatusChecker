package server

import (
	"github.com/Coderx44/StatusChecker/statuschecker"
)

type dependencies struct {
	httpchecker statuschecker.StatusChecker
}

func InitDependencies() (dependencies, error) {
	statuschecker := statuschecker.NewHttpChecker()

	return dependencies{
		httpchecker: statuschecker,
	}, nil
}
