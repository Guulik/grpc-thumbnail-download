package downloader

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDownloadRutube(t *testing.T) {
	tests := []struct {
		name    string
		videoId string
	}{
		{
			name:    "rutube",
			videoId: "fed7230fed49b9f6f5ae10922c275c93",
		},
		{
			name:    "utopia",
			videoId: "2d4f195278b9e1381e6cb983a114d389",
		},
	}
	for _, tt := range tests {

		ctx := context.Background()
		t.Run(tt.name, func(t *testing.T) {
			got, err := Download(ctx, tt.videoId)
			require.NoError(t, err)
			fmt.Println(got)
		})
	}
}

func TestDownloadYoutube(t *testing.T) {
	tests := []struct {
		name    string
		videoId string
	}{
		{
			name:    "redis",
			videoId: "QpBaA6B1U90",
		},
	}
	for _, tt := range tests {

		ctx := context.Background()
		t.Run(tt.name, func(t *testing.T) {
			got, err := Download(ctx, tt.videoId)
			require.NoError(t, err)
			fmt.Println(got)
		})
	}
}
