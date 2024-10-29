package IDextractor

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

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

func CheckRutube(videoURL string) bool {
	u, err := url.Parse(videoURL)
	if err != nil {
		fmt.Println("failed to parse url")
		return false
	}

	if u.Host == "rutube.ru" {
		fmt.Println("url is not from rutube")
		return true
	}
	return false
}

func validateURL(url string) error {
	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/(?:watch\?v=|embed\/|v\/|shorts\/)|youtu\.be\/)([a-zA-Z0-9_-]{11})(?:\S*)$`)
	matches := re.FindStringSubmatch(url)
	if len(matches) < 2 {
		return fmt.Errorf("%s: %s(%s)", "urlValidation", "URL is not valid", url)
	}
	return nil
}
