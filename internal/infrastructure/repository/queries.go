package repository

import (
	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func createEntityQuery(
	ds *goqu.SelectDataset,
	entity interface{},
) string {
	query, _, _ := ds.
		Insert().
		Rows(entity).
		ToSQL()
	return query
}

func getEntityByFieldQuery(
	ds *goqu.SelectDataset,
	fieldName, fieldVal string,
) string {
	query, _, _ := ds.
		Where(goqu.C(fieldName).Eq(fieldVal)).
		Limit(1).
		ToSQL()
	return query
}

func orderedExpression(params *domain.SelectParams) exp.OrderedExpression {
	orderedExpression := goqu.L(params.OrderBy).Asc()
	if params.Order == "desc" {
		orderedExpression = goqu.L(params.OrderBy).Desc()
	}
	return orderedExpression
}

func getEntitiesQuery(
	ds *goqu.SelectDataset,
	params *domain.SelectParams,
) string {
	query, _, _ := ds.
		Order(orderedExpression(params)).
		Limit(params.Limit).
		Offset(params.Offset).
		ToSQL()
	return query
}

func getEntitiesByFieldQuery(
	ds *goqu.SelectDataset,
	fieldName, fieldVal string,
	params *domain.SelectParams,
) string {
	query, _, _ := ds.
		Where(goqu.C(fieldName).Eq(fieldVal)).
		Order(orderedExpression(params)).
		Limit(params.Limit).
		Offset(params.Offset).
		ToSQL()
	return query
}

func searchEntitiesByFieldQuery(
	ds *goqu.SelectDataset,
	fieldName, fieldVal string,
	params *domain.SelectParams,
) string {
	query, _, _ := ds.
		Where(goqu.C(fieldName).Like("%" + fieldVal + "%")).
		Order(orderedExpression(params)).
		Limit(params.Limit).
		Offset(params.Offset).
		ToSQL()
	return query
}

// TODO: building queries with parameters ($1, $2, ...)???
func updateEntityFieldQuery(
	ds *goqu.SelectDataset,
	whereFieldName, whereFieldVal,
	updFieldName, updFieldVal string,
) string {
	query, _, _ := ds.
		Where(goqu.C(whereFieldName).Eq(whereFieldVal)).
		Update().Set(goqu.Record{updFieldName: updFieldVal}).
		ToSQL()
	return query
}

// TODO: error validation!
func deleteEntityQuery(
	ds *goqu.SelectDataset,
	fieldName,
	fieldVal string,
) string {
	query, _, _ := ds.
		Where(goqu.C(fieldName).Eq(fieldVal)).
		Delete().
		ToSQL()
	return query
}
