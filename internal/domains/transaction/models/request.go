package models

import outletModels "github.com/ibidathoillah/majoo-test/internal/domains/outlet/models"

type OmzetGrouBy string

const (
	OMZET_GROUP_BY_ALL    OmzetGrouBy = "merchant"
	OMZET_GROUP_BY_OUTLET OmzetGrouBy = "outlet"
)

type FindAllOmzetReport struct {
	UserID      *int64
	MerchantIDs []int64
	OutletIDs   []int64
	OutletHash  *map[int64]*outletModels.Outlet
	GroupBy     string `http_query:"group_by"`
	Month       int    `http_query:"month" validate:"required,gt=0,lt=13"`
	Year        int    `http_query:"year" validate:"required"`
	Page        int64  `http_query:"page"`
	Limit       int64  `http_query:"limit"`
	OrderBy     string `http_query:"orderby"`
	Search      string `http_query:"search"`
}
