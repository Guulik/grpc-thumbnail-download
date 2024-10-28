package IDextractor

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExtractId_Happy(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		expectedId string
	}{
		{
			name:       "grpc",
			url:        "https://www.youtube.com/watch?v=EURjTg5fw-E&t=4744s",
			expectedId: "EURjTg5fw-E",
		},
		{
			name:       "drobushevkii",
			url:        "https://www.youtube.com/watch?v=_JZDBCDMviw",
			expectedId: "_JZDBCDMviw",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := ExtractId(tt.url)
			require.NoError(t, err)
			require.Equal(t, tt.expectedId, id)
		})
	}
}

func TestExtractId_BadUrl(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		expectedId string
	}{
		{
			name:       "yandex search",
			url:        "https://yandex.ru/search/?clid=2358536&text=selectel&l10n=ru&lr=213",
			expectedId: "",
		},
		{
			name:       "gmail",
			url:        "https://mail.google.com/mail/u/1/#inbox?projector=1",
			expectedId: "",
		},
		{
			name:       "gmail with \"v\" key",
			url:        "https://mail.google.com/mail/u/1/#inbox?&v=_JZDBCDMviw",
			expectedId: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := ExtractId(tt.url)
			require.Errorf(t, err, "video ID not found in URL")
			require.Equal(t, tt.expectedId, id)
		})
	}
}

func TestExtractId_BadRandomSymbols(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		expectedId string
	}{
		{
			name:       "почему же",
			url:        "почему же твои губы так белеют на глазах почему же твои руки холодны в моих руках",
			expectedId: "",
		},
		{
			name:       "ящики",
			url:        "видишь ящики - грузи ящики",
			expectedId: "",
		},
		{
			name:       "??",
			url:        "Gaqwp;;фыафы13 -20яжэюдаd'f ;a ;la'",
			expectedId: "",
		},
		{
			name:       "есть ключ v",
			url:        "igrek?v=gaAGqet",
			expectedId: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := ExtractId(tt.url)
			require.Errorf(t, err, "URL is not valid")
			require.Equal(t, tt.expectedId, id)
		})
	}
}
