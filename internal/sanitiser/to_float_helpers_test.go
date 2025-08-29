package sanitiser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat(t *testing.T) {
	stringVal := "34.234"
	var exp float64 = 34.234
	result := ToFloat(stringVal)
	assert.Equal(t, exp, result)

	var val float64 = 34.34
	resultTwo := ToFloat(val)
	assert.Equal(t, val, resultTwo)
}
