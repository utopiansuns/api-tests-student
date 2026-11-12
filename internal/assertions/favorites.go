package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	favoritesModels "api-tests-template/internal/managers/favorites/models"
)

// FavoritesContainAdvertisement проверяет, что объявление есть в списке избранного
func FavoritesContainAdvertisement(t *testing.T, favorites favoritesModels.FavoritesList, advertisementId string) {
	t.Helper()
	require.Truef(t, containsFavorite(favorites, advertisementId),
		"объявления %s нет в избранном, а должно быть (всего в избранном: %d)", advertisementId, len(favorites.Items))
}

// FavoritesNotContainAdvertisement проверяет, что объявления нет в списке избранного
func FavoritesNotContainAdvertisement(t *testing.T, favorites favoritesModels.FavoritesList, advertisementId string) {
	t.Helper()
	require.Falsef(t, containsFavorite(favorites, advertisementId),
		"объявление %s есть в избранном, хотя должно было пропасть", advertisementId)
}

func containsFavorite(favorites favoritesModels.FavoritesList, advertisementId string) bool {
	for _, item := range favorites.Items {
		if item.Advertisement.Id == advertisementId {
			return true
		}
	}
	return false
}
