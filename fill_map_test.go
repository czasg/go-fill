package fill

import (
	"testing"
)

type FillMapTest struct {
	MapA map[string]string
	MapB map[string][]int
	MapC map[string]interface{}
}

func TestMapFill(t *testing.T) {
	assert := assertWrap(t)
	{
		base := FillMapTest{}
		assert("pass", Fill(&base), nil)
		assert("pass", base.MapA, map[string]string{})
		assert("pass", base.MapB, map[string][]int{})
		assert("pass", base.MapC, map[string]interface{}{})
	}
}
