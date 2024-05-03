package slug

import (
	"main/internal/assert"
	"testing"
)

func TestSlugify(t *testing.T) {
	assert.Equal("this-is-a-string", Slugify("ThIs Is a StrinG "))
	assert.Equal("remove-special-chars", Slugify("remove&Special^chars"))
}
