package favorites

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisement"
	favoriteManager "api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func TestFavoriteBusinessRules(t *testing.T) {
	base.SetupSuite()

	sellerToken, buyerToken := loginFavoriteUsers(t)

	cases := []struct {
		name          string
		token         string
		addBeforeTest bool
		wantStatus    int
		wantMessage   string
	}{
		{
			name:          "повторное добавление",
			token:         buyerToken,
			addBeforeTest: true,
			wantStatus:    http.StatusConflict,
			wantMessage:   "Advertisement is already in favorites",
		},
		{
			name:        "собственное объявление",
			token:       sellerToken,
			wantStatus:  http.StatusForbidden,
			wantMessage: "Cannot add own advertisement to favorites",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			created := createFavoriteTestAd(
				t, sellerToken, "Тест правил избранного", "Объявление для проверки правил избранного",
			)
			t.Cleanup(func() {
				t.Logf("Очистка: удаляем тестовое объявление %s", created.ID)
				advertisement.DeleteAdvertisement(t, sellerToken, created.ID)
			})

			if tc.addBeforeTest {
				t.Log("Покупатель впервые добавляет объявление в избранное")
				favoriteManager.AddFavorite(t, buyerToken, created.ID)
			}

			t.Logf("Проверяем случай %q, ожидаем HTTP %d", tc.name, tc.wantStatus)
			message := favoriteManager.AddFavoriteExpectError(
				t, tc.token, created.ID, tc.wantStatus,
			)
			t.Logf("Сообщение API: %s", message)
			require.Equal(t, tc.wantMessage, message,
				"Неверное сообщение API для случая %q", tc.name)
		})
	}
}
