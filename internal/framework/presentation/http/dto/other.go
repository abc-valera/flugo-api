package dto

import "github.com/abc-valera/flugo-api/internal/domain"

// SelectParamsQuery model info
//
//	@Description	Data for sorting
type SelectParamsQuery struct {
	OrderBy string `query:"order_by" json:"order_by" validate:"required" example:"created_at"`
	Order   string `query:"order"  json:"order" validate:"optional" enums:"asc,desc"`
	Limit   uint   `query:"limit"  json:"limit" validate:"required"`
	Offset  uint   `query:"offset" json:"offset"  validate:"required"`
}

func NewDomainSelectParams(params *SelectParamsQuery) *domain.SelectParams {
	return &domain.SelectParams{
		OrderBy: params.OrderBy,
		Order:   params.Order,
		Limit:   params.Limit,
		Offset:  params.Offset,
	}
}
