package favorites

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/helpers/fixtures"
	"api-tests-template/internal/managers/advertisements"
	advertisementsModels "api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/managers/auth"
	authModels "api-tests-template/internal/managers/auth/models"

	base "api-tests-template/tests"
)

// TestSuite использует двух пользователей: owner создаёт объявления (TEST_LOGIN из .env.override),
// buyer — новый пользователь, который добавляет их в избранное (своё объявление добавить нельзя)
type TestSuite struct {
	suite.Suite
	owner authModels.LoginOkResponse
	buyer authModels.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Авторизация владельца объявлений с кредами из переменных окружения")
	s.owner = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	base.Precondition("Регистрация нового пользователя-покупателя для работы с избранным")
	s.buyer = auth.RegisterRandomUser(s.T())
	require.NotEqual(s.T(), s.owner.User.Id, s.buyer.User.Id, "Владелец и покупатель должны быть разными пользователями")
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

// createOwnerAdvertisement создаёт объявление от имени владельца и регистрирует его удаление после теста
func (s *TestSuite) createOwnerAdvertisement() advertisementsModels.Advertisement {
	t := s.T()
	base.Precondition("Владелец создаёт тестовое объявление")
	advertisement := advertisements.CreateAdvertisement(t, s.owner.Token,
		fixtures.NewFullAdvertisementRequest(), fixtures.PhotoPaths(1))

	t.Cleanup(func() {
		base.Cleanup(fmt.Sprintf("Удаляем тестовое объявление %s", advertisement.Id))
		advertisements.DeleteAdvertisementIfExists(t, s.owner.Token, advertisement.Id)
	})

	return advertisement
}
