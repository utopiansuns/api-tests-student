package favorites

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	favoritesModels "api-tests-template/internal/managers/favorites/models"
)

// HttpGetFavorites получает страницу избранного авторизованного пользователя
func HttpGetFavorites(t *testing.T, token string, pageSize int) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Get(path.FavoritesPath).
		Query("page_size", strconv.Itoa(pageSize)).
		Expect(t).
		End().Response
}

// HttpPostFavorite добавляет объявление в избранное авторизованного пользователя
func HttpPostFavorite(t *testing.T, token string, request favoritesModels.AddFavoriteRequest) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Post(path.FavoritesPath).
		ContentType("application/json").
		JSON(request).
		Expect(t).
		End().Response
}

// HttpDeleteFavorite удаляет объявление из избранного авторизованного пользователя
func HttpDeleteFavorite(t *testing.T, token, advertisementId string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Delete(fmt.Sprintf(path.FavoriteByAdIdPath, advertisementId)).
		Expect(t).
		End().Response
}
