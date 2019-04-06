package http

import "github.com/nattaponra/iot-rule-engine/engine"

type apiReq interface {
	validate() error
}

type userReq struct {
	user engine.Rule
}

func (req userReq) validate() error {
	return req.user.Validate()
}
