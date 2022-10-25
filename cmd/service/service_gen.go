// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint "books/pkg/endpoint"
	http1 "books/pkg/http"
	service "books/pkg/service"
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"Create":                 {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Create", logger))},
		"Delete":                 {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Delete", logger))},
		"GetBook":                {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetBook", logger))},
		"Publish":                {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Publish", logger))},
		"SendEmailBookPublished": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SendEmailBookPublished", logger))},
		"Update":                 {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Update", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Create"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Create")), endpoint.InstrumentingMiddleware(duration.With("method", "Create"))}
	mw["Update"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Update")), endpoint.InstrumentingMiddleware(duration.With("method", "Update"))}
	mw["Delete"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Delete")), endpoint.InstrumentingMiddleware(duration.With("method", "Delete"))}
	mw["Publish"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Publish")), endpoint.InstrumentingMiddleware(duration.With("method", "Publish"))}
	mw["GetBook"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetBook")), endpoint.InstrumentingMiddleware(duration.With("method", "GetBook"))}
	mw["SendEmailBookPublished"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SendEmailBookPublished")), endpoint.InstrumentingMiddleware(duration.With("method", "SendEmailBookPublished"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Create", "Update", "Delete", "Publish", "GetBook", "SendEmailBookPublished"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
