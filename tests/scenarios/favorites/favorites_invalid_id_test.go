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

func TestAddFavoriteWithInvalidAdvertisementID(t *testing.T) {
	base.SetupSuite()

	t.Log("Входим под аккаунтом покупателя")
	buyer := auth.Login(t, os.Getenv("TEST_BUYER_LOGIN"), os.Getenv("TEST_BUYER_PASSWORD"))
	require.NotEmpty(t, buyer.Token, "В ответе на вход покупателя нет токена")

	const invalidAdID = "invalid-id"
	t.Logf("Проверяем отказ при добавлении объявления с некорректным ID %q", invalidAdID)
	message := favoriteManager.AddFavoriteExpectError(
		t, buyer.Token, invalidAdID, http.StatusBadRequest,
	)
	t.Logf("Сообщение API: %s", message)
	require.Equal(t, "invalid UUID length: 10", message,
		"API должно сообщить о неверной длине UUID объявления")
}
