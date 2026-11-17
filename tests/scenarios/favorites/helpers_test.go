package favorites

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/auth"

	"api-tests-template/internal/managers/advertisement"
	"api-tests-template/internal/utils"
	"path/filepath"
)

// loginFavoriteUsers входит под двумя разными аккаунтами и возвращает их токены.
func loginFavoriteUsers(t *testing.T) (sellerToken, buyerToken string) {
	t.Helper()

	t.Log("Входим под аккаунтами продавца и покупателя")
	seller := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	buyer := auth.Login(t, os.Getenv("TEST_BUYER_LOGIN"), os.Getenv("TEST_BUYER_PASSWORD"))

	require.NotEmpty(t, seller.Token, "В ответе на вход продавца нет токена")
	require.NotEmpty(t, buyer.Token, "В ответе на вход покупателя нет токена")
	require.NotEmpty(t, seller.User.Id, "В ответе на вход нет ID продавца")
	require.NotEmpty(t, buyer.User.Id, "В ответе на вход нет ID покупателя")
	require.NotEqual(t, seller.User.Id, buyer.User.Id, "Для теста нужны два разных аккаунта")

	return seller.Token, buyer.Token
}

// createFavoriteTestAd создаёт объявление для сценария избранного.
func createFavoriteTestAd(t *testing.T, sellerToken, titlePrefix, description string) advertisement.CreatedAdvertisement {
	t.Helper()

	fields := map[string]string{
		"title":       titlePrefix + " " + utils.RandomString(8),
		"description": description,
		"price":       "1557",
		"quantity":    "2",
	}
	photoPath := filepath.Join("..", "..", "testdata", "advertisement.png")

	t.Log("Продавец создаёт тестовое объявление")
	return advertisement.CreateAdvertisement(t, sellerToken, fields, photoPath)
}
