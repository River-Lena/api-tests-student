package advertisementCreate

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisement"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/utils"

	"api-tests-template/tests/testdata"

	base "api-tests-template/tests"
)

func TestCreateAdvertisementWithAllFields(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: входим под аккаунтом продавца")
	seller := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	require.NotEmpty(t, seller.Token, "В ответе на вход нет токена")
	require.NotEmpty(t, seller.User.Id, "В ответе на вход нет ID продавца")

	const (
		description = "Тестовое объявление для лабораторной"
		price       = testdata.AdvertisementPrice
		quantity    = testdata.AdvertisementQuantity
	)

	title := "Тестовая ручка " + utils.RandomString(8)
	photoPath := advertisementPhotoPath()
	fields := map[string]string{
		"title":       title,
		"description": description,
		"price":       strconv.Itoa(price),
		"quantity":    strconv.Itoa(quantity),
	}

	t.Log("Создаём объявление со всеми полями и одной фотографией")
	created := advertisement.CreateAdvertisement(t, seller.Token, fields, photoPath)
	t.Cleanup(func() {
		t.Logf("Очистка: удаляем объявление %s", created.ID)
		advertisement.DeleteAdvertisement(t, seller.Token, created.ID)
	})

	t.Log("Читаем объявление по ID и сверяем сохранённые поля")
	saved := advertisement.GetAdvertisement(t, seller.Token, created.ID)
	require.Equal(t, created.ID, saved.ID, "ID объявления не совпадает")
	require.Equal(t, title, saved.Title, "Название не совпадает")
	require.Equal(t, description, saved.Description, "Описание не совпадает")
	require.Equal(t, float64(price), saved.Price, "Цена не совпадает")
	require.Equal(t, quantity, saved.Quantity, "Количество не совпадает")
	require.Equal(t, seller.User.Id, saved.UserID, "ID владельца не совпадает")

	t.Log("Ищем объявление по полному названию")
	found := advList.GetAdvertisementsList(t, seller.Token, title)
	require.True(t, containsCreatedAdvertisement(found.Items, created.ID, title),
		"Поиск по %q не вернул объявление с ID %s", title, created.ID)

	t.Log("Проверяем фотографию через API")
	photos := advertisement.GetAdvertisementPhotos(t, seller.Token, created.ID)
	require.Len(t, photos, 1, "API должен вернуть одну загруженную фотографию")
	require.NotEmpty(t, photos[0].ID, "У фотографии нет ID")
	require.NotEmpty(t, photos[0].URL, "У фотографии нет URL")

	t.Log("Поля объявления, поиск и фотография проверены")
}

// containsCreatedAdvertisement проверяет, что поиск вернул созданное объявление.
func containsCreatedAdvertisement(items []models.Advertisement, id string, title string) bool {
	for _, item := range items {
		if item.ID == id && item.Title == title {
			return true
		}
	}
	return false
}
