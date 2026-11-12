package createAdvertisement

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/managers/auth"
	authModels "api-tests-template/internal/managers/auth/models"

	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
	loginData authModels.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Авторизация пользователя с кредами из переменных окружения")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}
