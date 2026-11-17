package favorites

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisement"
	favoriteManager "api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func TestFavoriteDisappearsAfterAdvertisementDeletion(t *testing.T) {
	base.SetupSuite()

	sellerToken, buyerToken := loginFavoriteUsers(t)

	created := createFavoriteTestAd(
		t, sellerToken, "Тест удаления", "Проверка избранного после удаления объявления",
	)

	deleted := false
	t.Cleanup(func() {
		if !deleted {
			t.Logf("Очистка: удаляем тестовое объявление %s", created.ID)
			advertisement.DeleteAdvertisement(t, sellerToken, created.ID)
		}
	})

	t.Log("Покупатель добавляет объявление и проверяет, что оно есть в избранном")
	favoriteManager.AddFavorite(t, buyerToken, created.ID)
	items := favoriteManager.GetFavorites(t, buyerToken)
	require.True(t, containsFavorite(items, created.ID),
		"Перед удалением объявление %q должно быть в избранном", created.ID)

	t.Log("Продавец удаляет объявление из каталога")
	advertisement.DeleteAdvertisement(t, sellerToken, created.ID)
	deleted = true

	t.Log("Покупатель снова запрашивает список избранного")
	items = favoriteManager.GetFavorites(t, buyerToken)
	require.False(t, containsFavorite(items, created.ID),
		"После удаления из каталога объявление %q осталось в избранном", created.ID)
}
