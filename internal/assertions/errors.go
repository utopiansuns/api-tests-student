package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	commonModels "api-tests-template/internal/managers/common/models"
)

// ErrorResponseEquals проверяет код и текст ошибки в ответе API
func ErrorResponseEquals(t *testing.T, expectedError, expectedMessage string, actual commonModels.ErrorResponse) {
	t.Helper()
	require.Equal(t, expectedError, actual.Error, "Некорректное значение error в ответе")
	require.Equal(t, expectedMessage, actual.Message, "Некорректное значение message в ответе")
}
