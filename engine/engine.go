package engine

import (
	"github.com/asaskevich/govalidator"
)

type Response interface {
	Code() int
	Headers() map[string]string
	Empty() bool
}

type Rule struct {
	Email    string
	Password string
}

func (u Rule) Validate() error {
	if u.Email == "" || u.Password == "" {
		return ErrMalformedEntity
	}

	if !govalidator.IsEmail(u.Email) {
		return ErrMalformedEntity
	}

	return nil
}

type RuleEngineRepository interface {
	Save(Rule) error
}
