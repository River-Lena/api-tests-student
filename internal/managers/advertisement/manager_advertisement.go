package advertisement

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	client "api-tests-template/internal/client/http/advertisement"
)

// CreatedAdvertisement описывает поля ответа, нужные для проверки создания.
type CreatedAdvertisement struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	UserID      string  `json:"user_id"`
}

// CreateAdvertisement создаёт объявление и возвращает данные из ответа API.
func CreateAdvertisement(
	t *testing.T,
	token string,
	fields map[string]string,
	photoPaths ...string,
) CreatedAdvertisement {
	t.Helper()

	contentType, body := buildAdvertisementForm(t, fields, photoPaths...)
	response := client.HttpPostAdvertisement(t, token, contentType, body)
	require.NotNil(t, response, "API не вернул HTTP-ответ")
	defer func() {
		if err := response.Body.Close(); err != nil {
			t.Errorf("Не удалось закрыть тело ответа: %v", err)
		}
	}()

	require.Equal(t, http.StatusCreated, response.StatusCode,
		"При создании объявления ожидали HTTP 201")

	var result CreatedAdvertisement
	err := json.NewDecoder(response.Body).Decode(&result)
	require.NoError(t, err, "Не удалось разобрать ответ о создании объявления")
	require.NotEmpty(t, result.ID, "В ответе нет ID созданного объявления")

	return result
}

// GetAdvertisement читает сохранённые данные объявления по ID.
func GetAdvertisement(t *testing.T, token string, id string) CreatedAdvertisement {
	t.Helper()

	response := client.HttpGetAdvertisement(t, token, id)
	require.NotNil(t, response, "API не вернул ответ для объявления %q", id)
	defer func() {
		if err := response.Body.Close(); err != nil {
			t.Errorf("Не удалось закрыть тело ответа: %v", err)
		}
	}()

	require.Equal(t, http.StatusOK, response.StatusCode,
		"При чтении объявления %q ожидали HTTP 200", id)

	var result CreatedAdvertisement
	err := json.NewDecoder(response.Body).Decode(&result)
	require.NoError(t, err, "Не удалось разобрать ответ для объявления %q", id)

	return result
}

// DeleteAdvertisement удаляет тестовое объявление по ID.
func DeleteAdvertisement(t *testing.T, token string, id string) {
	t.Helper()

	response := client.HttpDeleteAdvertisement(t, token, id)
	require.NotNil(t, response, "API не вернул ответ при удалении объявления %q", id)
	defer func() {
		if err := response.Body.Close(); err != nil {
			t.Errorf("Не удалось закрыть тело ответа: %v", err)
		}
	}()

	require.Equal(t, http.StatusNoContent, response.StatusCode,
		"При удалении объявления %q ожидали HTTP 204", id)
}

// AdvertisementPhoto содержит идентификатор и адрес загруженной фотографии.
type AdvertisementPhoto struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// GetAdvertisementPhotos читает список фотографий объявления по ID.
func GetAdvertisementPhotos(t *testing.T, token string, id string) []AdvertisementPhoto {
	t.Helper()

	response := client.HttpGetAdvertisementPhotos(t, token, id)
	require.NotNil(t, response, "API не вернул ответ со списком фото объявления %q", id)
	defer func() {
		if err := response.Body.Close(); err != nil {
			t.Errorf("Не удалось закрыть тело ответа: %v", err)
		}
	}()

	require.Equal(t, http.StatusOK, response.StatusCode,
		"При получении фото объявления %q ожидали HTTP 200", id)

	var photos []AdvertisementPhoto
	err := json.NewDecoder(response.Body).Decode(&photos)
	require.NoError(t, err, "Не удалось разобрать список фото объявления %q", id)

	return photos
}

// CreateAdvertisementExpectError проверяет статус ошибки и возвращает её сообщение.
func CreateAdvertisementExpectError(
	t *testing.T,
	token string,
	wantStatus int,
	fields map[string]string,
	photoPaths ...string,
) string {
	t.Helper()

	contentType, body := buildAdvertisementForm(t, fields, photoPaths...)
	response := client.HttpPostAdvertisement(t, token, contentType, body)
	require.NotNil(t, response, "API не вернул HTTP-ответ")
	defer func() {
		if err := response.Body.Close(); err != nil {
			t.Errorf("Не удалось закрыть тело ответа: %v", err)
		}
	}()

	require.Equal(t, wantStatus, response.StatusCode,
		"Неожиданный статус при создании объявления")

	var result struct {
		Message string `json:"message"`
	}
	err := json.NewDecoder(response.Body).Decode(&result)
	require.NoError(t, err, "Не удалось разобрать JSON ошибки")
	require.NotEmpty(t, result.Message, "В ответе API нет сообщения об ошибке")

	return result.Message
}
