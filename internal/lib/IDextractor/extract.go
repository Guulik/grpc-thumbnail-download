package IDextractor

import (
	"fmt"
	"net/url"
	"strings"
)

func ExtractId(videoURL string) (string, error) {
	isRutube := checkRutube(videoURL)
	var (
		videoId string
		err     error
	)
	if isRutube {
		videoId, err = ExtractIdRutube(videoURL)
	} else {
		videoId, err = ExtractIdYoutube(videoURL)
	}
	if err != nil {
		return "", err
	}
	return videoId, nil
}

func ExtractIdYoutube(videoURL string) (string, error) {
	/*	if err := validateURL(videoURL); err != nil {
		return "", err
	}*/

	u, err := url.Parse(videoURL)
	if err != nil {
		fmt.Println("failed to parse url")
		return "", fmt.Errorf("%s: %w", "ThumbnailService.ExtractIdYoutube", err)
	}
	query := u.Query()
	videoID := query.Get("v")
	if videoID == "" {
		fmt.Println("failed to get videoID")
		return "", fmt.Errorf("%s: %s", "ThumbnailService.ExtractIdYoutube", "video ID not found in URL")
	}
	return videoID, nil
}

func ExtractIdRutube(videoURL string) (string, error) {
	const op = "ThumbnailService.ExtractIdRutube"

	u, err := url.Parse(videoURL)
	if err != nil {
		fmt.Println("failed to parse url")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	pathSegments := strings.Split(u.Path, "/")
	if len(pathSegments) < 3 {
		fmt.Println("invalid path format")
		return "", fmt.Errorf("%s: %s", op, "invalid path format")
	}

	videoID := pathSegments[2] // предполагаем, что ID - 4-й сегмент
	if videoID == "" {
		fmt.Println("failed to get videoID")
		return "", fmt.Errorf("%s: %s", op, "video ID not found in URL")
	}

	return videoID, nil
}

func checkRutube(videoURL string) bool {
	u, err := url.Parse(videoURL)
	if err != nil {
		fmt.Println("failed to parse url")
		return false
	}

	if u.Host == "rutube.ru" {
		return true
	}
	return false
}
