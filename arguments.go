package gokore

type ListArguments struct {
	Offset    int
	Limit     int
	Sort      string
	Direction int
}

func NewListArguments(offset, limit int, sort string, direction int) *ListArguments {
	return &ListArguments{
		Offset:    offset,
		Limit:     limit,
		Sort:      sort,
		Direction: direction,
	}
}
