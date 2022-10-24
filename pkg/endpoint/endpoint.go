package endpoint

import (
	service "books/pkg/service"
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.BooksService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		rs, err := s.Create(ctx, req.Title, req.Author)
		return CreateResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.Err
}

// UpdateRequest collects the request parameters for the Update method.
type UpdateRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// UpdateResponse collects the response parameters for the Update method.
type UpdateResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUpdateEndpoint returns an endpoint that invokes Update on the service.
func MakeUpdateEndpoint(s service.BooksService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		rs, err := s.Update(ctx, req.Title, req.Author)
		return UpdateResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateResponse) Failed() error {
	return r.Err
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	ID int `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.BooksService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		rs, err := s.Delete(ctx, req.ID)
		return DeleteResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, title string, author string) (rs string, err error) {
	request := CreateRequest{
		Author: author,
		Title:  title,
	}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).Rs, response.(CreateResponse).Err
}

// Update implements Service. Primarily useful in a client.
func (e Endpoints) Update(ctx context.Context, title string, author string) (rs string, err error) {
	request := UpdateRequest{
		Author: author,
		Title:  title,
	}
	response, err := e.UpdateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateResponse).Rs, response.(UpdateResponse).Err
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id int) (rs string, err error) {
	request := DeleteRequest{
		ID: id,
	}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Rs, response.(DeleteResponse).Err
}

// PublishRequest collects the request parameters for the Publish method.
type PublishRequest struct {
	ID int `json:"id"`
}

// PublishResponse collects the response parameters for the Publish method.
type PublishResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakePublishEndpoint returns an endpoint that invokes Publish on the service.
func MakePublishEndpoint(s service.BooksService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PublishRequest)
		rs, err := s.Publish(ctx, req.ID)
		return PublishResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r PublishResponse) Failed() error {
	return r.Err
}

// Publish implements Service. Primarily useful in a client.
func (e Endpoints) Publish(ctx context.Context, id int) (rs string, err error) {
	request := PublishRequest{
		ID: id,
	}
	response, err := e.PublishEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PublishResponse).Rs, response.(PublishResponse).Err
}
