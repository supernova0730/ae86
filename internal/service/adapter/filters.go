package adapter

type ProductFilter struct {
	Title      *string
	MinPrice   *int
	MaxPrice   *int
	IsActive   *bool
	IsDeleted  *bool
	CategoryID *uint
}
