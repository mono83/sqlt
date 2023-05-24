package mysqlt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var nameData = []struct {
	Expected, Given string
}{
	{Expected: "`name`", Given: "name"},
	{Expected: "`name`", Given: "`name"},
	{Expected: "`name`", Given: "name`"},
	{Expected: "`name`", Given: "`name`"},
}

func TestName(t *testing.T) {
	for _, datum := range nameData {
		t.Run(fmt.Sprint(datum), func(t *testing.T) {
			assert.Equal(t, datum.Expected, Name(datum.Given))
		})
	}
}
