package service

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"thumbnail-proxy/internal/domain/model"
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

	log := t.log.With(op)

	var (
		err     error
		videoId string
		tb      model.Thumbnail
	)
	videoId, err = extractId(URL)
	if err != nil {
		log.Error("failed to extract videoID", sl.Err(err))
		return model.Thumbnail{}, fmt.Errorf("%s: %w", op, err)
	}

	tb, err = t.cacheTbProvider.Thumbnail(ctx, videoId)
	if err == nil {
		//TODO: maybe update in cache
		return tb, nil
	}

	var tbData []byte
	tbData, err = download(ctx, URL)

	thumbnail := model.Thumbnail{VideoId: videoId, Image: tbData}
	err = t.cacheTbSaver.SaveThumbnail(ctx, thumbnail, videoId)
	if err != nil {
		log.Error("failed to save videoID", sl.Err(err))
	}
	return thumbnail, nil
}

func download(ctx context.Context, url string) ([]byte, error) {
	//TODO: contextCancel timeout
	const op = "ThumbnailService.download"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to get thumbnail")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	// Проверяем, что запрос успешен
	if resp.StatusCode != http.StatusOK {
		fmt.Println("failed to get thumbnail")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read thumbnail image data")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return imageData, nil
}

func extractId(videoURL string) (string, error) {
	u, err := url.Parse(videoURL)
	if err != nil {
		fmt.Println("failed to parse url")
		return "", fmt.Errorf("%s: %w", "ThumbnailService.extractId", err)
	}
	query := u.Query()
	videoID := query.Get("v")
	if videoID == "" {
		fmt.Println("failed to get videoID")
		return "", fmt.Errorf("%s: %s", "ThumbnailService.extractId", "video ID not found in URL")
	}
	return videoID, nil
}
