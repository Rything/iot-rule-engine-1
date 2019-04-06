package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-zoo/bone"
	"github.com/nattaponra/iot-rule-engine/engine"
	log "github.com/nattaponra/iot-rule-engine/logger"
)

const contentType = "application/json"

var (
	errUnsupportedContentType = errors.New("unsupported content type")
	logger                    log.Logger
)

func MakeHandler(svc engine.Service, l log.Logger) http.Handler {
	logger = l

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	mux := bone.New()

	mux.Post("/rule", kithttp.NewServer(
		registrationEndpoint(svc),
		decodeCredentials,
		encodeResponse,
		opts...,
	))

	mux.Get("/rule", kithttp.NewServer(
		allRulesEndpoint(svc),
		decodeCredentials,
		encodeResponse,
		opts...,
	))

	return mux
}

func decodeCredentials(_ context.Context, r *http.Request) (interface{}, error) {
	if r.Header.Get("Content-Type") != contentType {
		logger.Warn("Invalid or missing content type.")
		return nil, errUnsupportedContentType
	}

	var user engine.Rule
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Warn(fmt.Sprintf("Failed to decode user credentials: %s", err))
		return nil, err
	}

	return userReq{user}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", contentType)

	if ar, ok := response.(engine.Response); ok {
		for k, v := range ar.Headers() {
			w.Header().Set(k, v)
		}

		w.WriteHeader(ar.Code())

		if ar.Empty() {
			return nil
		}
	}

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentType)

	switch err {
	case engine.ErrMalformedEntity:
		w.WriteHeader(http.StatusBadRequest)
	case engine.ErrUnauthorizedAccess:
		w.WriteHeader(http.StatusForbidden)
	case engine.ErrConflict:
		w.WriteHeader(http.StatusConflict)
	case errUnsupportedContentType:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	case io.ErrUnexpectedEOF:
		w.WriteHeader(http.StatusBadRequest)
	case io.EOF:
		w.WriteHeader(http.StatusBadRequest)
	default:
		switch err.(type) {
		case *json.SyntaxError:
			w.WriteHeader(http.StatusBadRequest)
		case *json.UnmarshalTypeError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
