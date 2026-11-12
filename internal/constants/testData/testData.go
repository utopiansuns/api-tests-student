package testData

const (
	// PhotoPath путь к тестовой фотографии относительно корня проекта
	PhotoPath = "testdata/photo.png"

	// InvalidToken заведомо невалидный авторизационный токен
	InvalidToken = "invalid_token"
	// EmptyToken отсутствующий токен: заголовок Authorization не отправляется
	EmptyToken = ""

	// TitleMaxLength максимальная длина названия объявления
	TitleMaxLength = 50
	// MaxPhotosCount максимальное количество фотографий в объявлении
	MaxPhotosCount = 3
	// MaxPrice максимальная цена объявления
	MaxPrice = 1000000
	// MaxQuantity максимальное количество товара в объявлении
	MaxQuantity = 100

	// AdvertisementTitlePrefix префикс названий тестовых объявлений
	AdvertisementTitlePrefix = "Autotest "
	// AdvertisementRandomPartLen длина случайной части названия, делает название уникальным
	AdvertisementRandomPartLen = 12
	// AdvertisementDescription описание тестового объявления
	AdvertisementDescription = "Объявление создано автотестом"
	// AdvertisementPrice цена тестового объявления
	AdvertisementPrice = 1500.5
	// AdvertisementQuantity количество товара в тестовом объявлении
	AdvertisementQuantity = 3
)
