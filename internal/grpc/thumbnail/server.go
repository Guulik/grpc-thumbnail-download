package thumbnail

import (
	"context"
	"google.golang.org/grpc"
	thumbnailv1 "thumbnail-proxy/proto/gen/thumbnail"
)

type Thumbnail interface {
	Get(
		ctx context.Context,
		url string,
	) (token string, err error)
}

type serverAPI struct {
	thumbnailv1.UnimplementedThumbnailServer
	thumbnail Thumbnail
}

func Register(gRPCServer *grpc.Server, thumbnail Thumbnail) {
	thumbnailv1.RegisterThumbnailServer(gRPCServer, &serverAPI{thumbnail: thumbnail})
}

func (s *serverAPI) Get(
	ctx context.Context,
	in *thumbnailv1.ThumbnailRequest,
) (*thumbnailv1.ThumbnailResponse, error) {
	//TODO: implement this
	panic("implement me")
}
