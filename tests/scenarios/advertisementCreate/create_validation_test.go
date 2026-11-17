package advertisementCreate

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"api-tests-template/internal/managers/advertisement"
	"api-tests-template/internal/managers/auth"

	"github.com/stretchr/testify/require"

	base "api-tests-template/tests"
)

func TestCreateAdvertisementWithoutTitle(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: входим под аккаунтом продавца")
	seller := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	fields := validAdvertisementFields("Тест валидации", "Проверка обязательного названия")
	delete(fields, "title")
	photoPath := advertisementPhotoPath()

	t.Log("Отправляем форму без поля title")
	message := advertisement.CreateAdvertisementExpectError(
		t, seller.Token, http.StatusBadRequest, fields, photoPath,
	)
	t.Logf("Сообщение API: %s", message)
	require.Equal(t, "title is required", message,
		"API должно сообщить, что название объявления обязательно")
}

func TestCreateAdvertisementInvalidValues(t *testing.T) {
	base.SetupSuite()
	t.Log("Входим под аккаунтом продавца")
	seller := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	require.NotEmpty(t, seller.Token, "В ответе на вход нет токена")

	photoPath := advertisementPhotoPath()
	cases := []struct {
		name        string
		field       string
		value       string
		wantMessage string
	}{
		{
			name:        "отрицательная цена",
			field:       "price",
			value:       "-1",
			wantMessage: "price must be between 0 and 1000000 inclusive",
		},
		{
			name:        "слишком длинное название",
			field:       "title",
			value:       strings.Repeat("A", 256),
			wantMessage: "title must not exceed 50 characters provided 256",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fields := validAdvertisementFields(
				"Тест валидации объявления",
				"Проверка некорректного значения",
			)
			fields[tc.field] = tc.value
			t.Logf("Проверяем поле %q", tc.field)
			message := advertisement.CreateAdvertisementExpectError(
				t, seller.Token, http.StatusBadRequest, fields, photoPath,
			)
			t.Logf("Сообщение API: %s", message)
			require.Equal(t, tc.wantMessage, message,
				"API вернуло неверное сообщение для поля %q", tc.field)
		})
	}
}

func TestCreateAdvertisementRequiredFields(t *testing.T) {
	base.SetupSuite()

	t.Log("Входим под аккаунтом продавца")
	seller := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	require.NotEmpty(t, seller.Token, "В ответе на вход нет токена")

	photoPath := advertisementPhotoPath()
	cases := []struct {
		name         string
		changeFields func(map[string]string)
		photos       []string
		wantMessage  string
	}{
		{
			name: "пустое название",
			changeFields: func(fields map[string]string) {
				fields["title"] = ""
			},
			photos:      []string{photoPath},
			wantMessage: "title is required",
		},
		{
			name: "без описания",
			changeFields: func(fields map[string]string) {
				delete(fields, "description")
			},
			photos:      []string{photoPath},
			wantMessage: "description is required",
		},
		{
			name:        "без фотографии",
			wantMessage: "at least 1 photo is required",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fields := validAdvertisementFields(
				"Тест обязательных полей",
				"Проверка валидации объявления",
			)
			if tc.changeFields != nil {
				tc.changeFields(fields)
			}

			t.Logf("Отправляем форму: %s", tc.name)
			message := advertisement.CreateAdvertisementExpectError(
				t, seller.Token, http.StatusBadRequest, fields, tc.photos...,
			)
			t.Logf("Сообщение API: %s", message)
			require.Equal(t, tc.wantMessage, message,
				"API вернуло неверное сообщение для случая %q", tc.name)
		})
	}
}
