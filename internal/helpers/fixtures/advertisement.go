package fixtures

import (
	"api-tests-template/internal/constants/testData"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/utils"
)

// NewFullAdvertisementRequest возвращает запрос на создание объявления со всеми параметрами
// и уникальным названием, чтобы тесты не пересекались между собой и между прогонами
func NewFullAdvertisementRequest() models.CreateAdvertisementRequest {
	return models.CreateAdvertisementRequest{
		Title:       testData.AdvertisementTitlePrefix + utils.RandomString(testData.AdvertisementRandomPartLen),
		Description: testData.AdvertisementDescription,
		Price:       testData.AdvertisementPrice,
		Quantity:    testData.AdvertisementQuantity,
	}
}

// PhotoPaths возвращает абсолютные пути к тестовой фотографии в указанном количестве
func PhotoPaths(count int) []string {
	paths := make([]string, 0, count)
	for range count {
		paths = append(paths, utils.ProjectFilePath(testData.PhotoPath))
	}
	return paths
}
