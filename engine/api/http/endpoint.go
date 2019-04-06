package http

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/nattaponra/iot-rule-engine/engine"
)

func allRulesEndpoint(svc engine.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		fmt.Println(request)
		return map[string]string{"message": "Hello"}, nil
	}
}

func registrationEndpoint(svc engine.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(userReq)

		if err := req.validate(); err != nil {
			return nil, err
		}

		err := svc.Create(req.user)
		return tokenRes{}, err
	}
}
