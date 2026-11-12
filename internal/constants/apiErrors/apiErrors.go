package apiErrors

// Коды ошибок (поле error в ответе API)
const (
	CodeInvalidRequest = "invalid_request"
	CodeUnauthorized   = "unauthorized"
	CodeForbidden      = "forbidden"
	CodeNotFound       = "not_found"
	CodeAlreadyExists  = "already_exists"
)

// Сообщения об ошибках (поле message в ответе API)
const (
	MessageAuthHeaderRequired      = "Authorization header required"
	MessageInvalidOrExpiredAuth    = "Invalid or expired token"
	MessageTitleRequired           = "title is required"
	MessageDescriptionRequired     = "description is required"
	MessagePhotoRequired           = "at least 1 photo is required"
	MessageTooManyPhotos           = "maximum 3 photos allowed"
	MessagePriceOutOfRange         = "price must be between 0 and 1000000 inclusive"
	MessageQuantityOutOfRange      = "quantity must be between 1 and 100 inclusive"
	MessageTitleTooLongFormat      = "title must not exceed %d characters provided %d"
	MessageAdvertisementNotFound   = "Advertisement not found"
	MessageFavoriteNotFound        = "Favorite entry not found"
	MessageFavoriteAlreadyExists   = "Advertisement is already in favorites"
	MessageCannotFavoriteOwnAd     = "Cannot add own advertisement to favorites"
	MessageInvalidAdvertisementId  = "Invalid advertisement ID"
	MessageInvalidUUIDLengthFormat = "invalid UUID length: %d"
)
