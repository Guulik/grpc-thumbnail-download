package cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"thumbnail-proxy/internal/lib/IDextractor"
	thumbnailv1 "thumbnail-proxy/proto/gen/thumbnail"
)

func (cli *CLI) getCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "getCommand [videoURL]",
		Short: "Get thumbnail from YouTube video URL",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			videoURL := args[0]
			return cli.get(ctx, videoURL)
		},
	}

}
func (cli *CLI) get(ctx context.Context, videoURL string) error {
	//TODO: contextCancel timeout
	const op = "cli.get"

	req := &thumbnailv1.ThumbnailRequest{URL: videoURL}
	resp, err := cli.client.GetThumbnail(ctx, req)
	if err != nil {
		fmt.Println("failed to get thumbnail from proxy")
		return fmt.Errorf("%s:%w", op, err)
	}

	videoID, err := IDextractor.ExtractId(videoURL)
	outputPath := filepath.Join(cli.cfg.OutputDir, fmt.Sprintf("%s_thumbnail.jpg", videoID))

	if err = os.MkdirAll(cli.cfg.OutputDir, os.ModePerm); err != nil {
		fmt.Println("failed to create output directory:")
		return fmt.Errorf("%s:%w", op, err)
	}

	if err = os.WriteFile(outputPath, resp.ThumbnailData, os.ModePerm); err != nil {
		fmt.Println("failed to save thumbnail to file")
		return fmt.Errorf("%s:%w", op, err)
	}

	fmt.Printf("Thumbnail saved to %s\n", outputPath)
	return nil
}
