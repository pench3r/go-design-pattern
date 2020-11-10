package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensitiveWordFilterChain(t *testing.T) {
	filterChain := &SensitiveWordFilterChain{}
	filterChain.AddFilter(&PoliticalWordFilter{})
	res := filterChain.Filter("hello china")
	assert.Equal(t, true, res)

	filterChain.AddFilter(&AdSensitiveWordFilter{})
	res = filterChain.Filter("11 of taobao")
	assert.Equal(t, true, res)
}
