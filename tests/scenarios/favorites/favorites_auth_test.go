package favorites

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	favoriteManager "api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func TestFavoritesWithoutToken(t *testing.T) {
	base.SetupSuite()

	const (
		token       = ""
		adID        = "00000000-0000-4000-8000-000000000000"
		wantMessage = "Authorization header required"
	)

	t.Run("получение списка", func(t *testing.T) {
		t.Log("Запрашиваем избранное без токена")
		message := favoriteManager.GetFavoritesExpectError(
			t, token, http.StatusUnauthorized,
		)
		require.Equal(t, wantMessage, message)
	})

	t.Run("добавление", func(t *testing.T) {
		t.Log("Добавляем объявление в избранное без токена")
		message := favoriteManager.AddFavoriteExpectError(
			t, token, adID, http.StatusUnauthorized,
		)
		require.Equal(t, wantMessage, message)
	})

	t.Run("удаление", func(t *testing.T) {
		t.Log("Удаляем объявление из избранного без токена")
		message := favoriteManager.DeleteFavoriteExpectError(
			t, token, adID, http.StatusUnauthorized,
		)
		require.Equal(t, wantMessage, message)
	})
}
