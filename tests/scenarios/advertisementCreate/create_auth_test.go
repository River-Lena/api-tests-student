package advertisementCreate

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisement"
	base "api-tests-template/tests"
)

func TestCreateAdvertisementUnauthorized(t *testing.T) {
	base.SetupSuite()

	photoPath := advertisementPhotoPath()
	cases := []struct {
		name        string
		token       string
		wantMessage string
	}{
		{
			name:        "без токена",
			token:       "",
			wantMessage: "Authorization header required",
		},
		{
			name:        "неверный токен",
			token:       "incorrect_token",
			wantMessage: "Invalid or expired token",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fields := validAdvertisementFields(
				"Тест авторизации",
				"Проверка создания объявления без действующего токена",
			)

			t.Logf("Отправляем корректную форму, случай %q", tc.name)
			message := advertisement.CreateAdvertisementExpectError(
				t, tc.token, http.StatusUnauthorized, fields, photoPath,
			)
			require.Equal(t, tc.wantMessage, message,
				"Неверное сообщение об ошибке для случая %q", tc.name)
			t.Logf("Сообщение API: %s", message)
		})
	}
}
