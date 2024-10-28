package redis

import (
	"context"
	"fmt"
	"log/slog"
	"thumbnail-proxy/internal/domain/model"
	"thumbnail-proxy/internal/lib/logger/sl"
	"thumbnail-proxy/internal/service"
)

var _ service.ThumbnailProvider = (*Cache)(nil)
var _ service.ThumbnailSaver = (*Cache)(nil)

func (c Cache) Thumbnail(ctx context.Context, videoId string) (model.Thumbnail, error) {
	const op = "Redis.Thumbnail"
	log := c.log.With(
		slog.String("op", op),
	)

	imageData, err := c.redis.Get(ctx, videoId).Bytes()
	if err != nil {
		log.Warn("failed to get cached thumbnail", sl.Err(err))
		return model.Thumbnail{}, fmt.Errorf("%s:%w", op, err)
	}

	return model.Thumbnail{VideoId: videoId, Image: imageData}, nil
}

func (c Cache) SaveThumbnail(ctx context.Context, thumbnail model.Thumbnail, videoId string) error {
	const op = "Redis.SaveThumbnail"
	log := c.log.With(
		slog.String("op", op),
	)

	imageData := thumbnail.Image
	err := c.redis.Set(ctx, videoId, imageData, c.cfg.Redis.TTL).Err()
	if err != nil {
		log.Error("failed to save thumbnail to cache", sl.Err(err))
		return fmt.Errorf("%s:%w", op, err)
	}
	return nil
}
