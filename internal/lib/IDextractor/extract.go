package IDextractor

import (
	"fmt"
	"net/url"
	"regexp"
)

func ExtractId(videoURL string) (string, error) {
	if err := validateURL(videoURL); err != nil {
		return "", err
	}

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

func validateURL(url string) error {
	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/(?:watch\?v=|embed\/|v\/|shorts\/)|youtu\.be\/)([a-zA-Z0-9_-]{11})(?:\S*)$`)
	matches := re.FindStringSubmatch(url)
	if len(matches) < 2 {
		return fmt.Errorf("%s: %s", "urlValidation", "URL is not valid")
	}
	return nil
}
