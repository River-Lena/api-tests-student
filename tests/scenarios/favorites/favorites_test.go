package favorites

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisement"
	favoriteManager "api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func TestAddAndRemoveFavorite(t *testing.T) {
	base.SetupSuite()

	sellerToken, buyerToken := loginFavoriteUsers(t)

	created := createFavoriteTestAd(
		t, sellerToken, "Тест избранного", "Объявление для проверки избранного",
	)

	t.Cleanup(func() {
		t.Logf("Очистка: удаляем тестовое объявление %s", created.ID)
		advertisement.DeleteAdvertisement(t, sellerToken, created.ID)
	})

	t.Log("Покупатель добавляет объявление в избранное")
	favoriteManager.AddFavorite(t, buyerToken, created.ID)
	items := favoriteManager.GetFavorites(t, buyerToken)
	require.True(t, containsFavorite(items, created.ID),
		"После добавления в избранном нет объявления %q", created.ID)

	t.Log("Покупатель удаляет объявление из избранного")
	favoriteManager.DeleteFavorite(t, buyerToken, created.ID)
	items = favoriteManager.GetFavorites(t, buyerToken)
	require.False(t, containsFavorite(items, created.ID),
		"После удаления объявление %q осталось в избранном", created.ID)

	t.Log("Появление и исчезновение объявления в списке проверены")
}

// containsFavorite проверяет наличие объявления по его ID в списке избранного.
func containsFavorite(items []favoriteManager.FavoriteItem, adID string) bool {
	for _, item := range items {
		if item.Advertisement.ID == adID {
			return true
		}
	}
	return false
}
