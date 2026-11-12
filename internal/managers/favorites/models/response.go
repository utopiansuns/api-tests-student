package models

import advertisementsModels "api-tests-template/internal/managers/advertisements/models"

type FavoritesList struct {
	Items    []FavoriteItem `json:"items"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	Total    int            `json:"total"`
}

type FavoriteItem struct {
	FavoriteId    string                             `json:"favorite_id"`
	AddedAt       string                             `json:"added_at"`
	Advertisement advertisementsModels.Advertisement `json:"advertisement"`
}
