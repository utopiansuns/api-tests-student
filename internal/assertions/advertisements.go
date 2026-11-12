package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisements/models"
)

// ContainsAdvertisementByTitle проверяет, что в списке есть объявление с точным названием
func ContainsAdvertisementByTitle(t *testing.T, items []models.Advertisement, title string) {
	t.Helper()
	for _, item := range items {
		if item.Title == title {
			return
		}
	}
	require.Failf(t, "объявление не найдено", "в ответе нет объявления с названием %q", title)
}

// AdvertisementMatchesRequest проверяет, что параметры объявления совпадают с переданными при создании
func AdvertisementMatchesRequest(t *testing.T, expected models.CreateAdvertisementRequest, actual models.Advertisement) {
	t.Helper()
	require.Equal(t, expected.Title, actual.Title, "title объявления не совпадает с переданным при создании")
	require.Equal(t, expected.Description, actual.Description, "description объявления не совпадает с переданным при создании")
	require.Equal(t, expected.Price, actual.Price, "price объявления не совпадает с переданным при создании")
	require.Equal(t, expected.Quantity, actual.Quantity, "quantity объявления не совпадает с переданным при создании")
}

// PhotosMatch проверяет, что список фотографий совпадает по id и url с ожидаемым
func PhotosMatch(t *testing.T, expected, actual []models.Photo) {
	t.Helper()
	require.Lenf(t, actual, len(expected), "ожидали %d фото, получили %d", len(expected), len(actual))
	for i := range expected {
		require.Equalf(t, expected[i].Id, actual[i].Id, "id фото #%d не совпадает", i)
		require.Equalf(t, expected[i].Url, actual[i].Url, "url фото #%d не совпадает", i)
	}
}
