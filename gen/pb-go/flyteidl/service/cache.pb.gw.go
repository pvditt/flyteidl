// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: flyteidl/service/cache.proto

/*
Package service is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package service

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_CacheService_EvictCache_0(ctx context.Context, marshaler runtime.Marshaler, client CacheServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EvictCacheRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["workflow_execution_id.project"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.project")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.project", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.project", err)
	}

	val, ok = pathParams["workflow_execution_id.domain"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.domain")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.domain", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.domain", err)
	}

	val, ok = pathParams["workflow_execution_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.name", err)
	}

	msg, err := client.EvictCache(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_CacheService_EvictCache_1 = &utilities.DoubleArray{Encoding: map[string]int{"task_execution_id": 0, "node_execution_id": 1, "execution_id": 2, "project": 3, "domain": 4, "name": 5, "node_id": 6, "task_id": 7, "version": 8, "retry_attempt": 9}, Base: []int{1, 20, 1, 1, 1, 4, 3, 2, 7, 5, 3, 0, 0, 0, 9, 6, 0, 15, 9, 0, 17, 12, 0, 19, 15, 0, 19, 18, 0, 20, 0}, Check: []int{0, 1, 2, 3, 4, 2, 6, 7, 2, 9, 10, 5, 8, 11, 2, 15, 16, 2, 18, 19, 2, 21, 22, 2, 24, 25, 2, 27, 28, 2, 30}}
)

func request_CacheService_EvictCache_1(ctx context.Context, marshaler runtime.Marshaler, client CacheServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EvictCacheRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["task_execution_id.node_execution_id.execution_id.project"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "task_execution_id.node_execution_id.execution_id.project")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "task_execution_id.node_execution_id.execution_id.project", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "task_execution_id.node_execution_id.execution_id.project", err)
	}

	val, ok = pathParams["task_execution_id.node_execution_id.execution_id.domain"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "task_execution_id.node_execution_id.execution_id.domain")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "task_execution_id.node_execution_id.execution_id.domain", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "task_execution_id.node_execution_id.execution_id.domain", err)
	}

	val, ok = pathParams["task_execution_id.node_execution_id.execution_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "task_execution_id.node_execution_id.execution_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "task_execution_id.node_execution_id.execution_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "task_execution_id.node_execution_id.execution_id.name", err)
	}

	val, ok = pathParams["task_execution_id.node_execution_id.node_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "task_execution_id.node_execution_id.node_id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "task_execution_id.node_execution_id.node_id", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "task_execution_id.node_execution_id.node_id", err)
	}

	val, ok = pathParams["task_execution_id.task_id.project"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "task_execution_id.task_id.project")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "task_execution_id.task_id.project", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "task_execution_id.task_id.project", err)
	}

	val, ok = pathParams["task_execution_id.task_id.domain"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "task_execution_id.task_id.domain")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "task_execution_id.task_id.domain", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "task_execution_id.task_id.domain", err)
	}

	val, ok = pathParams["task_execution_id.task_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "task_execution_id.task_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "task_execution_id.task_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "task_execution_id.task_id.name", err)
	}

	val, ok = pathParams["task_execution_id.task_id.version"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "task_execution_id.task_id.version")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "task_execution_id.task_id.version", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "task_execution_id.task_id.version", err)
	}

	val, ok = pathParams["task_execution_id.retry_attempt"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "task_execution_id.retry_attempt")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "task_execution_id.retry_attempt", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "task_execution_id.retry_attempt", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_CacheService_EvictCache_1); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.EvictCache(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterCacheServiceHandlerFromEndpoint is same as RegisterCacheServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCacheServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCacheServiceHandler(ctx, mux, conn)
}

// RegisterCacheServiceHandler registers the http handlers for service CacheService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCacheServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCacheServiceHandlerClient(ctx, mux, NewCacheServiceClient(conn))
}

// RegisterCacheServiceHandlerClient registers the http handlers for service CacheService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CacheServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CacheServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CacheServiceClient" to call the correct interceptors.
func RegisterCacheServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CacheServiceClient) error {

	mux.Handle("DELETE", pattern_CacheService_EvictCache_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CacheService_EvictCache_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CacheService_EvictCache_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_CacheService_EvictCache_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CacheService_EvictCache_1(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CacheService_EvictCache_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CacheService_EvictCache_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5, 1, 0, 4, 1, 5, 6}, []string{"api", "v1", "cache", "executions", "workflow_execution_id.project", "workflow_execution_id.domain", "workflow_execution_id.name"}, ""))

	pattern_CacheService_EvictCache_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5, 1, 0, 4, 1, 5, 6, 1, 0, 4, 1, 5, 7, 1, 0, 4, 1, 5, 8, 1, 0, 4, 1, 5, 9, 1, 0, 4, 1, 5, 10, 1, 0, 4, 1, 5, 11, 1, 0, 4, 1, 5, 12}, []string{"api", "v1", "cache", "task_executions", "task_execution_id.node_execution_id.execution_id.project", "task_execution_id.node_execution_id.execution_id.domain", "task_execution_id.node_execution_id.execution_id.name", "task_execution_id.node_execution_id.node_id", "task_execution_id.task_id.project", "task_execution_id.task_id.domain", "task_execution_id.task_id.name", "task_execution_id.task_id.version", "task_execution_id.retry_attempt"}, ""))
)

var (
	forward_CacheService_EvictCache_0 = runtime.ForwardResponseMessage

	forward_CacheService_EvictCache_1 = runtime.ForwardResponseMessage
)