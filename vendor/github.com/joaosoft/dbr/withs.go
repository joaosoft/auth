package dbr

import (
	"fmt"
)

type withs []*with

func newWith(name string, builder builder) *with {
	return &with{
		name:    name,
		builder: builder,
	}
}

type with struct {
	name    string
	builder builder
}

func (w withs) Build() (string, error) {

	if len(w) == 0 {
		return "", nil
	}

	var query string

	lenS := len(w)
	for i, item := range w {

		build, err := item.builder.Build()
		if err != nil {
			return "", err
		}

		query += fmt.Sprintf("%s %s (%s)", item.name, constFunctionAs, build)

		if i+1 < lenS {
			query += ", "
		}
	}

	return query, nil
}
