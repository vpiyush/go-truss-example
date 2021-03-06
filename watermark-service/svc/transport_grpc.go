// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: fcd9ff140d
// Version Date: 2021-07-14T06:36:40Z

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/vpiyush/go-kit-truss"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC WatermarkServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.WatermarkServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// watermark

		getdocument: grpctransport.NewServer(
			endpoints.GetDocumentEndpoint,
			DecodeGRPCGetDocumentRequest,
			EncodeGRPCGetDocumentResponse,
			serverOptions...,
		),
		adddocument: grpctransport.NewServer(
			endpoints.AddDocumentEndpoint,
			DecodeGRPCAddDocumentRequest,
			EncodeGRPCAddDocumentResponse,
			serverOptions...,
		),
		watermark: grpctransport.NewServer(
			endpoints.WatermarkEndpoint,
			DecodeGRPCWatermarkRequest,
			EncodeGRPCWatermarkResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the WatermarkServer interface
type grpcServer struct {
	getdocument grpctransport.Handler
	adddocument grpctransport.Handler
	watermark   grpctransport.Handler
}

// Methods for grpcServer to implement WatermarkServer interface

func (s *grpcServer) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (*pb.GetDocumentResponse, error) {
	_, rep, err := s.getdocument.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetDocumentResponse), nil
}

func (s *grpcServer) AddDocument(ctx context.Context, req *pb.AddDocumentRequest) (*pb.AddDocumentResponse, error) {
	_, rep, err := s.adddocument.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddDocumentResponse), nil
}

func (s *grpcServer) Watermark(ctx context.Context, req *pb.WatermarkRequest) (*pb.WatermarkResponse, error) {
	_, rep, err := s.watermark.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WatermarkResponse), nil
}

// Server Decode

// DecodeGRPCGetDocumentRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getdocument request to a user-domain getdocument request. Primarily useful in a server.
func DecodeGRPCGetDocumentRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetDocumentRequest)
	return req, nil
}

// DecodeGRPCAddDocumentRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC adddocument request to a user-domain adddocument request. Primarily useful in a server.
func DecodeGRPCAddDocumentRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddDocumentRequest)
	return req, nil
}

// DecodeGRPCWatermarkRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC watermark request to a user-domain watermark request. Primarily useful in a server.
func DecodeGRPCWatermarkRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.WatermarkRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCGetDocumentResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getdocument response to a gRPC getdocument reply. Primarily useful in a server.
func EncodeGRPCGetDocumentResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetDocumentResponse)
	return resp, nil
}

// EncodeGRPCAddDocumentResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain adddocument response to a gRPC adddocument reply. Primarily useful in a server.
func EncodeGRPCAddDocumentResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.AddDocumentResponse)
	return resp, nil
}

// EncodeGRPCWatermarkResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain watermark response to a gRPC watermark reply. Primarily useful in a server.
func EncodeGRPCWatermarkResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.WatermarkResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
