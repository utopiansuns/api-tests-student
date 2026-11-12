package models

import "strconv"

// CreateAdvertisementRequest параметры создания объявления
type CreateAdvertisementRequest struct {
	Title       string
	Description string
	Price       float64
	Quantity    int
}

// ToFormData превращает запрос в поля multipart-формы;
// title и description передаются всегда, price и quantity — только если заданы
func (r CreateAdvertisementRequest) ToFormData() map[string]string {
	formData := map[string]string{
		"title":       r.Title,
		"description": r.Description,
	}
	if r.Price != 0 {
		formData["price"] = strconv.FormatFloat(r.Price, 'f', -1, 64)
	}
	if r.Quantity != 0 {
		formData["quantity"] = strconv.Itoa(r.Quantity)
	}

	return formData
}
