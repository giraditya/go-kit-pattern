package http

import (
	endpoint "books/pkg/endpoint"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	http1 "github.com/go-kit/kit/transport/http"
)

// makeCreateHandler creates the handler logic
func makeCreateHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create", http1.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...))
}

// decodeCreateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateHandler creates the handler logic
func makeUpdateHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update", http1.NewServer(endpoints.UpdateEndpoint, decodeUpdateRequest, encodeUpdateResponse, options...))
}

// decodeUpdateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteHandler creates the handler logic
func makeDeleteHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete", http1.NewServer(endpoints.DeleteEndpoint, decodeDeleteRequest, encodeDeleteResponse, options...))
}

// decodeDeleteRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makePublishHandler creates the handler logic
func makePublishHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/publish", http1.NewServer(endpoints.PublishEndpoint, decodePublishRequest, encodePublishResponse, options...))
}

// decodePublishRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePublishRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.PublishRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePublishResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePublishResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
