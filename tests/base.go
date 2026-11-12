package tests

import (
	"log"

	"api-tests-template/internal/utils"
)

// SetupSuite действия перед запуском сьюта с тестами
func SetupSuite() {
	log.Println("Init environment variables")
	utils.LoadEnv()
}

// TearDownSuite действия после запуска сьюта с тестами
func TearDownSuite() {
	log.Println("Tear down suite")
}

// Precondition выводит текст в консоль с меткой [Precondition]
func Precondition(text string) {
	utils.LogWithLabelAndTimeStamp("Precondition", text)
}

// Step выводит в консоль описание шага теста с меткой [Step]
func Step(text string) {
	utils.LogWithLabelAndTimeStamp("Step", text)
}

// Check выводит в консоль описание проверки с меткой [Check]
func Check(text string) {
	utils.LogWithLabelAndTimeStamp("Check", text)
}

// Cleanup выводит в консоль описание действия по очистке тестовых данных с меткой [Cleanup]
func Cleanup(text string) {
	utils.LogWithLabelAndTimeStamp("Cleanup", text)
}
