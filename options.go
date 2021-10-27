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

type ListChannelByThingOptions struct {
	offset       int
	limit        int
	order        ThingOrder
	direction    Direction
	disconnected bool
}

func NewListChannelByThingOptions() *ListChannelByThingOptions {
	return &ListChannelByThingOptions{
		offset:       0,
		limit:        10,
		order:        THING_ORDER_NAME,
		direction:    DIRECTION_DESC,
		disconnected: true,
	}
}

func (opts *ListChannelByThingOptions) SetOffset(offset int) *ListChannelByThingOptions {
	opts.offset = offset
	return opts
}

func (opts *ListChannelByThingOptions) SetLimit(limit int) *ListChannelByThingOptions {
	opts.limit = limit
	return opts
}

func (opts *ListChannelByThingOptions) SetOrder(order ThingOrder) *ListChannelByThingOptions {
	opts.order = order
	return opts
}

func (opts *ListChannelByThingOptions) SetDirection(dir Direction) *ListChannelByThingOptions {
	opts.direction = dir
	return opts
}

func (opts *ListChannelByThingOptions) SetDisconnected(disconnected bool) *ListChannelByThingOptions {
	opts.disconnected = disconnected
	return opts
}

type ListAllChannelOptions struct {
	offset    int
	limit     int
	order     ThingOrder
	direction Direction
}

func NewListAllChannelOptions() *ListAllChannelOptions {
	return &ListAllChannelOptions{
		offset:    0,
		limit:     10,
		order:     THING_ORDER_NAME,
		direction: DIRECTION_DESC,
	}
}

func (opts *ListAllChannelOptions) SetOffset(offset int) *ListAllChannelOptions {
	opts.offset = offset
	return opts
}

func (opts *ListAllChannelOptions) SetLimit(limit int) *ListAllChannelOptions {
	opts.limit = limit
	return opts
}

func (opts *ListAllChannelOptions) SetOrder(order ThingOrder) *ListAllChannelOptions {
	opts.order = order
	return opts
}

func (opts *ListAllChannelOptions) SetDirection(dir Direction) *ListAllChannelOptions {
	opts.direction = dir
	return opts
}

type ListChannelByUserOptions struct {
	offset    int
	limit     int
	order     ThingOrder
	direction Direction
}

func NewListChannelByUserOptions() *ListChannelByUserOptions {
	return &ListChannelByUserOptions{
		offset:    0,
		limit:     10,
		order:     THING_ORDER_NAME,
		direction: DIRECTION_DESC,
	}
}

func (opts *ListChannelByUserOptions) SetOffset(offset int) *ListChannelByUserOptions {
	opts.offset = offset
	return opts
}

func (opts *ListChannelByUserOptions) SetLimit(limit int) *ListChannelByUserOptions {
	opts.limit = limit
	return opts
}

func (opts *ListChannelByUserOptions) SetOrder(order ThingOrder) *ListChannelByUserOptions {
	opts.order = order
	return opts
}

func (opts *ListChannelByUserOptions) SetDirection(dir Direction) *ListChannelByUserOptions {
	opts.direction = dir
	return opts
}
