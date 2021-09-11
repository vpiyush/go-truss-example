package handlers

import (
	"context"

	pb "github.com/vpiyush/go-kit-truss"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.WatermarkServer {
	return watermarkService{}
}

type watermarkService struct {
	//repo svc.Repository
}

func (s watermarkService) GetDocument(ctx context.Context, in *pb.GetDocumentRequest) (*pb.GetDocumentResponse, error) {
	var resp pb.GetDocumentResponse
	return &resp, nil
}

func (s watermarkService) AddDocument(ctx context.Context, in *pb.AddDocumentRequest) (*pb.AddDocumentResponse, error) {
	var resp pb.AddDocumentResponse
	return &resp, nil
}

func (s watermarkService) Watermark(ctx context.Context, in *pb.WatermarkRequest) (*pb.WatermarkResponse, error) {
	var resp pb.WatermarkResponse
	return &resp, nil
}
