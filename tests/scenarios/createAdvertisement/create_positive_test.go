package createAdvertisement

import (
	"fmt"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testData"
	"api-tests-template/internal/helpers/fixtures"
	"api-tests-template/internal/managers/advertisements"
	advertisementsModels "api-tests-template/internal/managers/advertisements/models"

	base "api-tests-template/tests"
)

// Positive: создание объявления со всеми параметрами, проверка карточки, поиска и фотографий
func (s *TestSuite) TestCreateAdvertisementWithAllParams() {
	request := fixtures.NewFullAdvertisementRequest()
	photoPaths := fixtures.PhotoPaths(testData.MaxPhotosCount)

	var created advertisementsModels.Advertisement
	s.Run("Создаём объявление со всеми параметрами", func() {
		base.Step(fmt.Sprintf("Создаём объявление %q с %d фото", request.Title, len(photoPaths)))
		created = advertisements.CreateAdvertisement(s.T(), s.loginData.Token, request, photoPaths)
		base.Check(fmt.Sprintf("Объявление создано, id=%s", created.Id))
		require.NotEmpty(s.T(), created.Id, "В ответе на создание нет id объявления")
		assertions.AdvertisementMatchesRequest(s.T(), request, created)
		require.Equal(s.T(), s.loginData.User.Id, created.UserId, "Объявление создано не от нашего пользователя")
	})
	s.T().Cleanup(func() {
		if created.Id == "" {
			return
		}
		base.Cleanup(fmt.Sprintf("Удаляем тестовое объявление %s", created.Id))
		advertisements.DeleteAdvertisementIfExists(s.T(), s.loginData.Token, created.Id)
	})

	s.Run("GET /advertisement возвращает те же параметры, что передали при создании", func() {
		base.Step("Получаем объявление по id")
		advertisement := advertisements.GetAdvertisementById(s.T(), created.Id)
		base.Check("Сравниваем параметры объявления с переданными при создании")
		require.Equal(s.T(), created.Id, advertisement.Id, "id объявления не совпадает")
		assertions.AdvertisementMatchesRequest(s.T(), request, advertisement)
	})

	s.Run("Объявление находится в поиске по полному названию", func() {
		base.Step(fmt.Sprintf("Ищем объявление по названию %q", request.Title))
		list := advertisements.GetAdvertisementsList(s.T(), s.loginData.Token, request.Title)
		base.Check("Проверяем, что объявление есть в результатах поиска")
		assertions.ContainsAdvertisementByTitle(s.T(), list.Items, request.Title)
	})

	s.Run("Фотографии загружены и доступны через GET /advertisements/{id}/photos", func() {
		base.Step("Получаем фотографии объявления")
		photos := advertisements.GetAdvertisementPhotos(s.T(), s.loginData.Token, created.Id)
		base.Check(fmt.Sprintf("Проверяем, что загружено %d фото и они совпадают с ответом на создание", len(photoPaths)))
		assertions.PhotosMatch(s.T(), created.Photos, photos)
		for _, photo := range photos {
			require.NotEmpty(s.T(), photo.Url, "У фотографии пустой url")
		}
	})
}
