package thumbnail

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"thumbnail-proxy/internal/domain/model"
	thumbnailv1 "thumbnail-proxy/proto/gen/thumbnail"
)

type Thumbnail interface {
	GetThumbnail(
		ctx context.Context,
		url string,
	) (model.Thumbnail, error)
}

type serverAPI struct {
	thumbnailv1.UnimplementedThumbnailServer
	service Thumbnail
}

func Register(gRPCServer *grpc.Server, service Thumbnail) {
	thumbnailv1.RegisterThumbnailServer(gRPCServer, &serverAPI{service: service})
}

func (s *serverAPI) GetThumbnail(
	ctx context.Context,
	in *thumbnailv1.ThumbnailRequest,
) (*thumbnailv1.ThumbnailResponse, error) {
	tb, err := s.service.GetThumbnail(ctx, in.GetURL())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get thumbnail")
	}
	return &thumbnailv1.ThumbnailResponse{ThumbnailData: tb.Image}, nil
}
