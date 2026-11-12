package favorites

import (
	"fmt"
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apiErrors"
	"api-tests-template/internal/constants/testData"
	commonModels "api-tests-template/internal/managers/common/models"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/utils"

	base "api-tests-template/tests"
)

// invalidAdvertisementId id объявления, который не является UUID
const invalidAdvertisementId = "abc"

// Negative: добавление и удаление из избранного несуществующего объявления возвращает 404
func (s *TestSuite) TestFavoritesNonExistentAdvertisement() {
	nonExistentId := utils.RandomUUID()

	s.Run("Добавление несуществующего объявления", func() {
		base.Step(fmt.Sprintf("Добавляем в избранное несуществующее объявление %s", nonExistentId))
		errorResponse := favorites.AddFavoriteWithError(s.T(), s.buyer.Token, nonExistentId, http.StatusNotFound)

		base.Check("Проверяем ошибку 404")
		assertions.ErrorResponseEquals(s.T(), apiErrors.CodeNotFound, apiErrors.MessageAdvertisementNotFound, errorResponse)
	})

	s.Run("Удаление несуществующего объявления", func() {
		base.Step(fmt.Sprintf("Удаляем из избранного несуществующее объявление %s", nonExistentId))
		errorResponse := favorites.RemoveFavoriteWithError(s.T(), s.buyer.Token, nonExistentId, http.StatusNotFound)

		base.Check("Проверяем ошибку 404")
		assertions.ErrorResponseEquals(s.T(), apiErrors.CodeNotFound, apiErrors.MessageFavoriteNotFound, errorResponse)
	})
}

// Negative: id объявления в неверном формате возвращает 400
func (s *TestSuite) TestFavoritesInvalidAdvertisementId() {
	s.Run("Добавление с id не в формате UUID", func() {
		base.Step(fmt.Sprintf("Добавляем в избранное объявление с id %q", invalidAdvertisementId))
		errorResponse := favorites.AddFavoriteWithError(s.T(), s.buyer.Token, invalidAdvertisementId, http.StatusBadRequest)

		base.Check("Проверяем ошибку 400")
		assertions.ErrorResponseEquals(s.T(), apiErrors.CodeInvalidRequest,
			fmt.Sprintf(apiErrors.MessageInvalidUUIDLengthFormat, len(invalidAdvertisementId)), errorResponse)
	})

	s.Run("Удаление с id не в формате UUID", func() {
		base.Step(fmt.Sprintf("Удаляем из избранного объявление с id %q", invalidAdvertisementId))
		errorResponse := favorites.RemoveFavoriteWithError(s.T(), s.buyer.Token, invalidAdvertisementId, http.StatusBadRequest)

		base.Check("Проверяем ошибку 400")
		assertions.ErrorResponseEquals(s.T(), apiErrors.CodeInvalidRequest, apiErrors.MessageInvalidAdvertisementId, errorResponse)
	})
}

// Negative: ручки избранного без авторизации или с неверным токеном возвращают 401
func (s *TestSuite) TestFavoritesUnauthorized() {
	advertisementId := utils.RandomUUID()

	tokens := []struct {
		name            string
		token           string
		expectedMessage string
	}{
		{"без токена", testData.EmptyToken, apiErrors.MessageAuthHeaderRequired},
		{"неверный токен", testData.InvalidToken, apiErrors.MessageInvalidOrExpiredAuth},
	}

	requests := []struct {
		name string
		call func(t *testing.T, token string) commonModels.ErrorResponse
	}{
		{"GET /favorites", func(t *testing.T, token string) commonModels.ErrorResponse {
			return favorites.GetFavoritesWithError(t, token, http.StatusUnauthorized)
		}},
		{"POST /favorites", func(t *testing.T, token string) commonModels.ErrorResponse {
			return favorites.AddFavoriteWithError(t, token, advertisementId, http.StatusUnauthorized)
		}},
		{"DELETE /favorites/{ad_id}", func(t *testing.T, token string) commonModels.ErrorResponse {
			return favorites.RemoveFavoriteWithError(t, token, advertisementId, http.StatusUnauthorized)
		}},
	}

	for _, request := range requests {
		for _, tc := range tokens {
			s.Run(fmt.Sprintf("%s %s", request.name, tc.name), func() {
				base.Step(fmt.Sprintf("Вызываем %s: %s", request.name, tc.name))
				errorResponse := request.call(s.T(), tc.token)

				base.Check("Проверяем ошибку 401")
				assertions.ErrorResponseEquals(s.T(), apiErrors.CodeUnauthorized, tc.expectedMessage, errorResponse)
			})
		}
	}
}

// Negative: повторное добавление того же объявления в избранное возвращает 409
func (s *TestSuite) TestFavoritesAddDuplicate() {
	advertisement := s.createOwnerAdvertisement()

	base.Step("Покупатель добавляет объявление в избранное первый раз")
	favorites.AddFavorite(s.T(), s.buyer.Token, advertisement.Id)

	base.Step("Покупатель добавляет то же объявление повторно")
	errorResponse := favorites.AddFavoriteWithError(s.T(), s.buyer.Token, advertisement.Id, http.StatusConflict)

	base.Check("Проверяем ошибку 409")
	assertions.ErrorResponseEquals(s.T(), apiErrors.CodeAlreadyExists, apiErrors.MessageFavoriteAlreadyExists, errorResponse)
}

// Negative: добавление собственного объявления в избранное возвращает 403
func (s *TestSuite) TestFavoritesAddOwnAdvertisement() {
	advertisement := s.createOwnerAdvertisement()

	base.Step("Владелец пытается добавить своё объявление в избранное")
	errorResponse := favorites.AddFavoriteWithError(s.T(), s.owner.Token, advertisement.Id, http.StatusForbidden)

	base.Check("Проверяем ошибку 403")
	assertions.ErrorResponseEquals(s.T(), apiErrors.CodeForbidden, apiErrors.MessageCannotFavoriteOwnAd, errorResponse)
}
