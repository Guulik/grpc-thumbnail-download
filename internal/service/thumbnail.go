package service

import (
	"context"
	"log/slog"
	"thumbnail-proxy/internal/domain/model"
	"thumbnail-proxy/internal/lib/IDextractor"
	"thumbnail-proxy/internal/lib/downloader"
	"thumbnail-proxy/internal/lib/logger/sl"
	"time"
)

type ThumbnailService struct {
	log             *slog.Logger
	cacheTbProvider ThumbnailProvider
	cacheTbSaver    ThumbnailSaver
	timeout         time.Duration
}

type ThumbnailProvider interface {
	Thumbnail(ctx context.Context, videoId string) (model.Thumbnail, error)
}

type ThumbnailSaver interface {
	SaveThumbnail(ctx context.Context, thumbnail model.Thumbnail, videoId string) error
}

func New(
	log *slog.Logger,
	provider ThumbnailProvider,
	saver ThumbnailSaver,
	timeout time.Duration,
) *ThumbnailService {
	return &ThumbnailService{
		log:             log,
		cacheTbProvider: provider,
		cacheTbSaver:    saver,
		timeout:         timeout,
	}
}

func (t *ThumbnailService) GetThumbnail(ctx context.Context, URL string) (model.Thumbnail, error) {
	//TODO: contextCancel timeout
	const op = "Service.GetThumbnail"

	log := t.log.With(slog.String("op", op))

	var (
		err     error
		videoId string
		tb      model.Thumbnail
	)

	videoId, err = IDextractor.ExtractId(URL)
	tb, err = t.cacheTbProvider.Thumbnail(ctx, videoId)
	if err == nil {
		//TODO: maybe update in cache
		return tb, nil
	}

	var tbData []byte
	log.Info("Trying to download img from video: ", URL)
	tbData, err = downloader.Download(ctx, videoId)
	if err != nil {
		return model.Thumbnail{}, err
	}
	log.Info("image data:", tbData)

	thumbnail := model.Thumbnail{VideoId: videoId, Image: tbData}
	err = t.cacheTbSaver.SaveThumbnail(ctx, thumbnail, videoId)
	if err != nil {
		log.Warn("failed to save thumbnail", sl.Err(err))
	}
	return thumbnail, nil
}
