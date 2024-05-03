package sanitiser

import (
	"main/internal/assert"
	"testing"
)

func TestToFloat(t *testing.T) {
	stringVal := "34.234"
	var exp float64 = 34.234
	result := ToFloat(stringVal)
	assert.Equal(exp, result)

	var val float64 = 34.34
	resultTwo := ToFloat(val)
	assert.Equal(val, resultTwo)
}
