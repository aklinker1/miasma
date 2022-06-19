package server

type Pagination struct {
	Page int32
	Size int32
}

func (p Pagination) Limit() int32 {
	return p.Size
}

func (p Pagination) Offset() int32 {
	zeroIndexPage := p.Page - 1
	if zeroIndexPage < 0 {
		zeroIndexPage = 0
	}
	return zeroIndexPage * p.Size
}
