package engine

import "errors"

var (
	ErrConflict           = errors.New("email already taken")
	ErrMalformedEntity    = errors.New("malformed entity specification")
	ErrUnauthorizedAccess = errors.New("missing or invalid credentials provided")
	ErrNotFound           = errors.New("non-existent entity")
)

type Service interface {
	Create(Rule) error
}

var _ Service = (*ruleEngineService)(nil)

type ruleEngineService struct {
	ruleEngine RuleEngineRepository
}

func New(ruleEnRepo RuleEngineRepository) Service {
	return &ruleEngineService{ruleEngine: ruleEnRepo}
}

func (r ruleEngineService) Create(rule Rule) error {
	return r.ruleEngine.Save(rule)
}
