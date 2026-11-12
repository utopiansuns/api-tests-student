package advertisements

import (
	"fmt"
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

const photosFormField = "photos"

// HttpGetAdvertisements ищет объявления по подстроке в названии, ожидает 200
func HttpGetAdvertisements(t *testing.T, token, textForSearch string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Get(path.AdvertisementsPath).
		Query("search", textForSearch).
		Expect(t).
		Status(http.StatusOK).
		End().Response
}

// HttpPostAdvertisement создаёт объявление через multipart/form-data: текстовые поля и фотографии
func HttpPostAdvertisement(t *testing.T, token string, formFields map[string]string, photoPaths []string) *http.Response {
	request := apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath)
	for name, value := range formFields {
		request = request.MultipartFormData(name, value)
	}
	if len(photoPaths) > 0 {
		request = request.MultipartFile(photosFormField, photoPaths...)
	}

	return request.Expect(t).End().Response
}

// HttpGetAdvertisementById получает объявление по его id
func HttpGetAdvertisementById(t *testing.T, id string) *http.Response {
	return apiRunner.GetRunner().Create().
		Get(path.AdvertisementPath).
		Query("id", id).
		Expect(t).
		End().Response
}

// HttpDeleteAdvertisement удаляет объявление по его id
func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Delete(path.AdvertisementPath).
		Query("id", id).
		Expect(t).
		End().Response
}

// HttpGetAdvertisementPhotos получает список фотографий объявления
func HttpGetAdvertisementPhotos(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Get(fmt.Sprintf(path.AdvertisementPhotosPath, id)).
		Expect(t).
		End().Response
}
