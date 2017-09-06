package services

import "github.com/datake914/lonbutu/middleware"

type ServiceExecuter interface {
	Execute()
}

type serviceExecuter struct {
	service Service
}

func NewServiceExecuter(service Service) ServiceExecuter {
	return &serviceExecuter{service}
}

func (s *serviceExecuter) Execute() {
	tx := middleware.NewDatabase().Begin()
	s.service.Execute(tx)
	tx.Commit()
}
