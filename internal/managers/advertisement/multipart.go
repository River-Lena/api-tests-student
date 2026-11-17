package advertisement

import (
	"bytes"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// buildAdvertisementForm собирает поля и фотографии для создания объявления.
func buildAdvertisementForm(
	t *testing.T,
	fields map[string]string,
	photoPaths ...string,
) (contentType string, body string) {
	t.Helper()

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	for name, value := range fields {
		err := writer.WriteField(name, value)
		require.NoError(t, err, "Не удалось записать поле %q", name)
	}

	for _, photoPath := range photoPaths {
		data, err := os.ReadFile(photoPath)
		require.NoError(t, err, "Не удалось прочитать фото %q", photoPath)

		part, err := writer.CreateFormFile("photos", filepath.Base(photoPath))
		require.NoError(t, err, "Не удалось добавить фото %q в форму", photoPath)

		_, err = part.Write(data)
		require.NoError(t, err, "Не удалось записать фото %q в форму", photoPath)
	}

	require.NoError(t, writer.Close(), "Не удалось завершить форму")
	return writer.FormDataContentType(), buffer.String()
}
