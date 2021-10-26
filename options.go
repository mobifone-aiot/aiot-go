package aiot

type Direction string
type ThingOrder string

var (
	DIRECTION_ASC  Direction = "asc"
	DIRECTION_DESC Direction = "desc"

	THING_ORDER_NAME ThingOrder = "name"
	THING_ORDER_KEY  ThingOrder = "key"
	THING_ORDER_ID   ThingOrder = "id"
)

type ListThingsByUserOptions struct {
	offset    int
	limit     int
	order     ThingOrder
	direction Direction
}

func NewListThingsByUserOptions() *ListThingsByUserOptions {
	return &ListThingsByUserOptions{
		offset:    0,
		limit:     10,
		order:     THING_ORDER_NAME,
		direction: DIRECTION_DESC,
	}
}

func (opts *ListThingsByUserOptions) SetOffset(offset int) *ListThingsByUserOptions {
	opts.offset = offset
	return opts
}

func (opts *ListThingsByUserOptions) SetLimit(limit int) *ListThingsByUserOptions {
	opts.limit = limit
	return opts
}

func (opts *ListThingsByUserOptions) SetOrder(order ThingOrder) *ListThingsByUserOptions {
	opts.order = order
	return opts
}

func (opts *ListThingsByUserOptions) SetDirection(dir Direction) *ListThingsByUserOptions {
	opts.direction = dir
	return opts
}
