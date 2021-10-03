package query

type Builder struct {
	text string
	args []interface{}
}

func NewBuilder() *Builder {
	return &Builder{}
}
