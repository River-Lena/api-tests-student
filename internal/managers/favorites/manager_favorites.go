package favorites

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	client "api-tests-template/internal/client/http/favorites"
)

// FavoriteItem описывает запись избранного с вложенным объявлением.
type FavoriteItem struct {
	Advertisement struct {
		ID string `json:"id"`
	} `json:"advertisement"`
}

// AddFavorite добавляет объявление в избранное и проверяет HTTP 201.
func AddFavorite(t *testing.T, token string, adID string) {
	t.Helper()

	response := client.HttpPostFavorite(t, token, adID)
	require.NotNil(t, response, "API не вернул ответ при добавлении %q", adID)
	defer closeResponseBody(t, response)

	require.Equal(t, http.StatusCreated, response.StatusCode,
		"При добавлении объявления %q в избранное ожидали HTTP 201", adID)
}

// GetFavorites читает первую страницу избранного пользователя.
func GetFavorites(t *testing.T, token string) []FavoriteItem {
	t.Helper()

	response := client.HttpGetFavorites(t, token)
	require.NotNil(t, response, "API не вернул список избранного")
	defer closeResponseBody(t, response)

	require.Equal(t, http.StatusOK, response.StatusCode,
		"При чтении избранного ожидали HTTP 200")

	var result struct {
		Items []FavoriteItem `json:"items"`
	}
	err := json.NewDecoder(response.Body).Decode(&result)
	require.NoError(t, err, "Не удалось разобрать список избранного")

	return result.Items
}

// DeleteFavorite удаляет объявление из избранного и проверяет HTTP 204.
func DeleteFavorite(t *testing.T, token string, adID string) {
	t.Helper()

	response := client.HttpDeleteFavorite(t, token, adID)
	require.NotNil(t, response, "API не вернул ответ при удалении %q", adID)
	defer closeResponseBody(t, response)

	require.Equal(t, http.StatusNoContent, response.StatusCode,
		"При удалении объявления %q из избранного ожидали HTTP 204", adID)
}

// closeResponseBody закрывает тело HTTP-ответа и сообщает об ошибке закрытия.
func closeResponseBody(t *testing.T, response *http.Response) {
	t.Helper()

	if err := response.Body.Close(); err != nil {
		t.Errorf("Не удалось закрыть тело ответа: %v", err)
	}
}

// GetFavoritesExpectError проверяет ошибку при чтении избранного.
func GetFavoritesExpectError(t *testing.T, token string, wantStatus int) string {
	t.Helper()
	return readFavoriteError(t, client.HttpGetFavorites(t, token), wantStatus)
}

// AddFavoriteExpectError проверяет ошибку при добавлении в избранное.
func AddFavoriteExpectError(t *testing.T, token string, adID string, wantStatus int) string {
	t.Helper()
	return readFavoriteError(t, client.HttpPostFavorite(t, token, adID), wantStatus)
}

// DeleteFavoriteExpectError проверяет ошибку при удалении из избранного.
func DeleteFavoriteExpectError(t *testing.T, token string, adID string, wantStatus int) string {
	t.Helper()
	return readFavoriteError(t, client.HttpDeleteFavorite(t, token, adID), wantStatus)
}

// readFavoriteError проверяет статус и читает сообщение из ответа API.
func readFavoriteError(t *testing.T, response *http.Response, wantStatus int) string {
	t.Helper()

	require.NotNil(t, response, "API избранного не вернул HTTP-ответ")
	defer closeResponseBody(t, response)

	require.Equal(t, wantStatus, response.StatusCode,
		"Неожиданный HTTP-статус для операции с избранным")

	var result struct {
		Message string `json:"message"`
	}
	err := json.NewDecoder(response.Body).Decode(&result)
	require.NoError(t, err, "Не удалось разобрать JSON ошибки избранного")
	require.NotEmpty(t, result.Message, "Ответ API не содержит сообщения об ошибке")

	return result.Message
}
