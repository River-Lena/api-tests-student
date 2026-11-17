package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpGetFavorites получает список избранного пользователя.
func HttpGetFavorites(t *testing.T, token string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Get(path.FavoritesPath).
		Expect(t).
		End().Response
}

// HttpPostFavorite добавляет объявление в избранное.
func HttpPostFavorite(t *testing.T, token string, adID string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Post(path.FavoritesPath).
		ContentType("application/json").
		JSON(map[string]string{"advertisement_id": adID}).
		Expect(t).
		End().Response
}

// HttpDeleteFavorite удаляет объявление из избранного по ID объявления.
func HttpDeleteFavorite(t *testing.T, token string, adID string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Delete(path.FavoritesPath + "/" + adID).
		Expect(t).
		End().Response
}
