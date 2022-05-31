package internal

type Pagination struct {
	Page int32
	Size int32
}

func (p Pagination) Limit() int32 {
	return p.Size
}

func (p Pagination) Offset() int32 {
	return (p.Page - 1) * p.Size
}

type Sort struct {
	Field     string
	Direction string
}

type AppsFilter struct {
	ID            *string
	Name          *string
	NameContains  *string
	IncludeHidden *bool
	Sort          *Sort
	Pagination    *Pagination
}
