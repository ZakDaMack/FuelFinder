package slug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlugify(t *testing.T) {
	assert.Equal(t, "this-is-a-string", Slugify("ThIs Is a StrinG "))
	assert.Equal(t, "remove-special-chars", Slugify("remove&Special^chars"))
}
