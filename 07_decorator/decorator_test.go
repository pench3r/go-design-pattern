package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	cs := NewColorSquare("blue", Square{})
	assert.Equal(t, "ColorSquare: this is square with blue", cs.Draw())
}
