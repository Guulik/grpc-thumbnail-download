package cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"sync"
	"thumbnail-proxy/internal/lib/IDextractor"
	thumbnailv1 "thumbnail-proxy/proto/gen/thumbnail"
)

func (cli *CLI) getCommand(ctx context.Context) *cobra.Command {
	var async bool
	cmd := &cobra.Command{
		Use:   "get [videoURLs...]",
		Short: "Get thumbnail from YouTube/Rutube video URL",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.get(ctx, args, async)
		},
	}
	cmd.Flags().BoolVar(&async, "async", false, "Enable asynchronous downloading")
	return cmd
}

func (cli *CLI) get(ctx context.Context, videoURLs []string, async bool) error {
	if async {
		return cli.getAsync(ctx, videoURLs)
	} else {
		return cli.getSync(ctx, videoURLs)
	}
}

// Синхронное скачивание
func (cli *CLI) getSync(ctx context.Context, videoURLs []string) error {
	for _, videoURL := range videoURLs {
		if err := cli.getRequest(ctx, videoURL); err != nil {
			return err
		}
	}
	return nil
}

// Асинхронное скачивание
func (cli *CLI) getAsync(ctx context.Context, videoURLs []string) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(videoURLs))

	for _, videoURL := range videoURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if err := cli.getRequest(ctx, url); err != nil {
				errCh <- err
			}
		}(videoURL)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}
	return nil
}

func (cli *CLI) getRequest(ctx context.Context, videoURL string) error {
	const op = "cli.getRequest"

	req := &thumbnailv1.ThumbnailRequest{URL: videoURL}
	resp, err := cli.client.GetThumbnail(ctx, req)
	if err != nil {
		fmt.Println("failed to getRequest thumbnail from proxy")
		return fmt.Errorf("%s:%w", op, err)
	}

	videoId, err := IDextractor.ExtractId(videoURL)
	err = cli.saveThumbnail(videoId, resp.ThumbnailData)
	return nil
}

func (cli *CLI) saveThumbnail(videoID string, imageData []byte) error {
	const op = "cli.saveThumbnail"

	outputPath := filepath.Join(cli.cfg.OutputDir, fmt.Sprintf("%s_thumbnail.jpg", videoID))
	fmt.Println("outputPath:", cli.cfg.OutputDir)

	if err := os.MkdirAll(cli.cfg.OutputDir, os.ModePerm); err != nil {
		fmt.Println("failed to create output directory:")
		return fmt.Errorf("%s:%w", op, err)
	}

	if err := os.WriteFile(outputPath, imageData, os.ModePerm); err != nil {
		fmt.Println("failed to save thumbnail to file")
		return fmt.Errorf("%s:%w", op, err)
	}

	fmt.Printf("Thumbnail saved to %s\n", outputPath)

	return nil
}
