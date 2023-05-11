package domain

// SelectParams represents query data for specifying select details.
// OrderField is field by which sorting will be performed (usually is 'created_at'),
// Order is order of sorting ('acs' or 'desc'),
// Limit limits number of returned units,
// Offset sets an offset for returned units.
type SelectParams struct {
	OrderBy string
	Order   string
	Limit   uint
	Offset  uint
}

// TODO: add validation
