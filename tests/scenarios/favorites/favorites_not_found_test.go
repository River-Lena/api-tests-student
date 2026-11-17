package favorites

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/auth"
	favoriteManager "api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func TestFavoriteAdvertisementNotFound(t *testing.T) {
	base.SetupSuite()

	t.Log("Входим под аккаунтом покупателя")
	buyer := auth.Login(t, os.Getenv("TEST_BUYER_LOGIN"), os.Getenv("TEST_BUYER_PASSWORD"))
	require.NotEmpty(t, buyer.Token, "В ответе на вход покупателя нет токена")

	const missingAdID = "00000000-0000-4000-8000-000000000000"

	t.Run("добавление несуществующего объявления", func(t *testing.T) {
		t.Logf("Добавляем в избранное объявление %s", missingAdID)
		message := favoriteManager.AddFavoriteExpectError(
			t, buyer.Token, missingAdID, http.StatusNotFound,
		)
		t.Logf("Сообщение API: %s", message)
		require.Equal(t, "Advertisement not found", message,
			"При добавлении API должно сообщить, что объявление не найдено")
	})

	t.Run("удаление несуществующего объявления", func(t *testing.T) {
		t.Logf("Удаляем из избранного объявление %s", missingAdID)
		message := favoriteManager.DeleteFavoriteExpectError(
			t, buyer.Token, missingAdID, http.StatusNotFound,
		)
		t.Logf("Сообщение API: %s", message)
		require.Equal(t, "Favorite entry not found", message,
			"При удалении API должно сообщить, что запись в избранном не найдена")
	})
}
