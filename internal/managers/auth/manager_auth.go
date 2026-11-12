package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
)

const (
	randomUserEmailDomain = "example.com"
	randomUserPasswordLen = 16
	randomUserSuffixLen   = 10
	randomUserFirstName   = "Autotest"
)

// Login выполняет авторизацию пользователя и возвращает токен с данными пользователя, ожидает 200
func Login(t *testing.T, username, password string) models.LoginOkResponse {
	require.NotEmpty(t, username, "Отсутствует логин пользователя")
	require.NotEmpty(t, password, "Отсутствует пароль пользователя")

	request := models.LoginRequest{
		Email:    username,
		Password: password,
	}

	loginRawResponse := auth.HttpPostAuthLogin(t, request)
	defer loginRawResponse.Body.Close()

	respArray, errResponse := io.ReadAll(loginRawResponse.Body)
	require.NoError(t, errResponse)

	loginResponse := models.LoginOkResponse{}
	err := json.Unmarshal(respArray, &loginResponse)
	require.NoError(t, err)

	return loginResponse
}

// RegisterRandomUser регистрирует нового пользователя со случайными email и паролем
// и возвращает его токен; нужен, когда в тесте требуется второй пользователь
func RegisterRandomUser(t *testing.T) models.LoginOkResponse {
	request := models.RegisterRequest{
		Email:     fmt.Sprintf("autotest_%s@%s", strings.ToLower(utils.RandomString(randomUserSuffixLen)), randomUserEmailDomain),
		Password:  utils.RandomString(randomUserPasswordLen),
		FirstName: randomUserFirstName,
	}

	registerRawResponse := auth.HttpPostAuthRegister(t, request)
	require.NoError(t, registerRawResponse.Body.Close())

	return Login(t, request.Email, request.Password)
}
