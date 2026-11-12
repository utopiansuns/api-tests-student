package advertisements

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/managers/common"
	commonModels "api-tests-template/internal/managers/common/models"
)

// GetAdvertisementsList ищет объявления по тексту и возвращает типизированный список
func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	raw := advertisements.HttpGetAdvertisements(t, token, text)
	defer raw.Body.Close()

	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err)

	resp := models.AdvertisementsList{}
	require.NoError(t, json.Unmarshal(body, &resp))
	return resp
}

// CreateAdvertisement создаёт объявление и возвращает его, ожидает 201
func CreateAdvertisement(t *testing.T, token string, request models.CreateAdvertisementRequest, photoPaths []string) models.Advertisement {
	t.Helper()
	raw := advertisements.HttpPostAdvertisement(t, token, request.ToFormData(), photoPaths)

	advertisement := models.Advertisement{}
	common.DecodeBody(t, raw, http.StatusCreated, &advertisement)
	return advertisement
}

// CreateAdvertisementWithError пытается создать объявление и возвращает тело ошибки, ожидает переданный статус-код
func CreateAdvertisementWithError(t *testing.T, token string, request models.CreateAdvertisementRequest,
	photoPaths []string, expectedStatusCode int) commonModels.ErrorResponse {
	t.Helper()
	raw := advertisements.HttpPostAdvertisement(t, token, request.ToFormData(), photoPaths)
	return common.DecodeError(t, raw, expectedStatusCode)
}

// GetAdvertisementById получает объявление по id, ожидает 200
func GetAdvertisementById(t *testing.T, id string) models.Advertisement {
	t.Helper()
	raw := advertisements.HttpGetAdvertisementById(t, id)

	advertisement := models.Advertisement{}
	common.DecodeBody(t, raw, http.StatusOK, &advertisement)
	return advertisement
}

// GetAdvertisementPhotos получает фотографии объявления, ожидает 200
func GetAdvertisementPhotos(t *testing.T, token, id string) []models.Photo {
	t.Helper()
	raw := advertisements.HttpGetAdvertisementPhotos(t, token, id)

	var photos []models.Photo
	common.DecodeBody(t, raw, http.StatusOK, &photos)
	return photos
}

// DeleteAdvertisement удаляет объявление, ожидает 204
func DeleteAdvertisement(t *testing.T, token, id string) {
	t.Helper()
	raw := advertisements.HttpDeleteAdvertisement(t, token, id)
	common.ReadBodyWithStatus(t, raw, http.StatusNoContent)
}

// DeleteAdvertisementIfExists удаляет объявление в cleanup-е теста, если оно ещё существует;
// объявление могло быть удалено в самом тесте, а повторное удаление сервер отклоняет
func DeleteAdvertisementIfExists(t *testing.T, token, id string) {
	t.Helper()
	getResponse := advertisements.HttpGetAdvertisementById(t, id)
	require.NoError(t, getResponse.Body.Close())
	if getResponse.StatusCode == http.StatusNotFound {
		return
	}

	DeleteAdvertisement(t, token, id)
}
