package common

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/common/models"
)

// ReadBodyWithStatus проверяет статус-код ответа и возвращает его тело;
// при несовпадении статуса в сообщении об ошибке выводится тело ответа
func ReadBodyWithStatus(t *testing.T, response *http.Response, expectedStatusCode int) []byte {
	t.Helper()
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err, "Не удалось прочитать тело ответа")

	require.Equalf(t, expectedStatusCode, response.StatusCode,
		"Ожидали HTTP %d, получили %d. Тело ответа: %s", expectedStatusCode, response.StatusCode, string(body))

	return body
}

// DecodeBody проверяет статус-код ответа и декодирует JSON-тело в переданную структуру
func DecodeBody(t *testing.T, response *http.Response, expectedStatusCode int, target any) {
	t.Helper()
	body := ReadBodyWithStatus(t, response, expectedStatusCode)
	require.NoErrorf(t, json.Unmarshal(body, target), "Не удалось распарсить JSON ответа: %s", string(body))
}

// DecodeError проверяет статус-код ответа и возвращает тело ошибки
func DecodeError(t *testing.T, response *http.Response, expectedStatusCode int) models.ErrorResponse {
	t.Helper()
	errorResponse := models.ErrorResponse{}
	DecodeBody(t, response, expectedStatusCode, &errorResponse)
	return errorResponse
}
