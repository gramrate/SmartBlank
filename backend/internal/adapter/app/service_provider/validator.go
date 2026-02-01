package service_provider

import "backend/internal/adapter/controller/validator"

func (s *ServiceProvider) Validator() *validator.Validator {
	if s.validator == nil {
		s.validator = validator.New()
	}
	return s.validator
}
