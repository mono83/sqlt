package mosaic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectChanges(t *testing.T) {
	ex := []existing{
		{ID: 9, Data: []any{"foo", 77}},
		{ID: 32, Data: []any{"bar", false}},
		{ID: 77, Data: []any{"baz", 77}},
		{ID: 93, Data: []any{"any", 2}},
	}

	data := [][]any{
		{"bar", true},
		{"baz", 77},
		{"new", 12},
	}

	update, del, insert := detectChanges(ex, data)
	assert.Equal(t, []uint64{77}, update)
	assert.Equal(t, []uint64{9, 32, 93}, del)
	assert.Equal(t, [][]any{{"bar", true}, {"new", 12}}, insert)
}
