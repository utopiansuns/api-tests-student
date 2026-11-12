package createAdvertisement

import (
	"fmt"
	"net/http"
	"strings"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apiErrors"
	"api-tests-template/internal/constants/testData"
	"api-tests-template/internal/helpers/fixtures"
	"api-tests-template/internal/managers/advertisements"
	advertisementsModels "api-tests-template/internal/managers/advertisements/models"

	base "api-tests-template/tests"
)

// Negative: создание объявления без обязательных или с некорректными параметрами возвращает 400
func (s *TestSuite) TestCreateAdvertisementInvalidParams() {
	tooLongTitleLength := testData.TitleMaxLength + 1

	cases := []struct {
		name            string
		modify          func(r *advertisementsModels.CreateAdvertisementRequest)
		photosCount     int
		expectedMessage string
	}{
		{
			name:            "пустой title",
			modify:          func(r *advertisementsModels.CreateAdvertisementRequest) { r.Title = "" },
			photosCount:     1,
			expectedMessage: apiErrors.MessageTitleRequired,
		},
		{
			name:            "пустой description",
			modify:          func(r *advertisementsModels.CreateAdvertisementRequest) { r.Description = "" },
			photosCount:     1,
			expectedMessage: apiErrors.MessageDescriptionRequired,
		},
		{
			name:            "без фотографий",
			modify:          func(_ *advertisementsModels.CreateAdvertisementRequest) {},
			photosCount:     0,
			expectedMessage: apiErrors.MessagePhotoRequired,
		},
		{
			name:            "фотографий больше максимума",
			modify:          func(_ *advertisementsModels.CreateAdvertisementRequest) {},
			photosCount:     testData.MaxPhotosCount + 1,
			expectedMessage: apiErrors.MessageTooManyPhotos,
		},
		{
			name: "title длиннее максимума",
			modify: func(r *advertisementsModels.CreateAdvertisementRequest) {
				r.Title = strings.Repeat("a", tooLongTitleLength)
			},
			photosCount:     1,
			expectedMessage: fmt.Sprintf(apiErrors.MessageTitleTooLongFormat, testData.TitleMaxLength, tooLongTitleLength),
		},
		{
			name:            "отрицательная цена",
			modify:          func(r *advertisementsModels.CreateAdvertisementRequest) { r.Price = -1 },
			photosCount:     1,
			expectedMessage: apiErrors.MessagePriceOutOfRange,
		},
		{
			name:            "цена больше максимума",
			modify:          func(r *advertisementsModels.CreateAdvertisementRequest) { r.Price = testData.MaxPrice + 1 },
			photosCount:     1,
			expectedMessage: apiErrors.MessagePriceOutOfRange,
		},
		{
			name:            "отрицательное количество",
			modify:          func(r *advertisementsModels.CreateAdvertisementRequest) { r.Quantity = -1 },
			photosCount:     1,
			expectedMessage: apiErrors.MessageQuantityOutOfRange,
		},
		{
			name:            "количество больше максимума",
			modify:          func(r *advertisementsModels.CreateAdvertisementRequest) { r.Quantity = testData.MaxQuantity + 1 },
			photosCount:     1,
			expectedMessage: apiErrors.MessageQuantityOutOfRange,
		},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			request := fixtures.NewFullAdvertisementRequest()
			tc.modify(&request)

			base.Step(fmt.Sprintf("Пытаемся создать объявление: %s", tc.name))
			errorResponse := advertisements.CreateAdvertisementWithError(s.T(), s.loginData.Token, request,
				fixtures.PhotoPaths(tc.photosCount), http.StatusBadRequest)

			base.Check(fmt.Sprintf("Проверяем ошибку: %q", tc.expectedMessage))
			assertions.ErrorResponseEquals(s.T(), apiErrors.CodeInvalidRequest, tc.expectedMessage, errorResponse)
		})
	}
}

// Negative: создание объявления без авторизации или с неверным токеном возвращает 401
func (s *TestSuite) TestCreateAdvertisementUnauthorized() {
	cases := []struct {
		name            string
		token           string
		expectedMessage string
	}{
		{"без токена", testData.EmptyToken, apiErrors.MessageAuthHeaderRequired},
		{"неверный токен", testData.InvalidToken, apiErrors.MessageInvalidOrExpiredAuth},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			base.Step(fmt.Sprintf("Пытаемся создать объявление: %s", tc.name))
			errorResponse := advertisements.CreateAdvertisementWithError(s.T(), tc.token,
				fixtures.NewFullAdvertisementRequest(), fixtures.PhotoPaths(1), http.StatusUnauthorized)

			base.Check("Проверяем, что вернулась ошибка 401 с корректным сообщением")
			assertions.ErrorResponseEquals(s.T(), apiErrors.CodeUnauthorized, tc.expectedMessage, errorResponse)
		})
	}
}
