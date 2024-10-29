package downloader

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"thumbnail-proxy/internal/lib/logger/handlers/slogpretty"
	"thumbnail-proxy/internal/lib/logger/sl"
)

func Download(ctx context.Context, videoId string) ([]byte, error) {
	//TODO: contextCancel timeout
	const op = "Downloader.download"

	log := setupPrettySlog()
	log = log.With(slog.String("op", op))

	var url string
	if len(videoId) > 15 {
		//У рутуба ID роликов состоит из 30 символов
		url = fmt.Sprintf("https://rutube.ru/api/video/%s/thumbnail/?redirect=1", videoId)
	} else {
		//У youtube ID роликов состоит из 11 символов
		url = fmt.Sprintf("https://img.youtube.com/vi/%s/maxresdefault.jpg", videoId)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}

	log.Info("trying to get img by url", slog.String("url", url))
	resp, err := client.Do(req)
	if err != nil {
		log.Error("failed to connect to youtube", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	log.Info("Status:", slog.String("status code", resp.Status))
	if resp.StatusCode != http.StatusOK {
		log.Warn("status code is not 200: ", slog.String("status code", strconv.Itoa(resp.StatusCode)))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read thumbnail image data")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return imageData, nil
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
