package fill

import (
	"testing"
)

type FillSliceTest struct {
	SliceA []int
	SliceB []map[string]string
	SliceC []interface{}
}

func TestSliceFill(t *testing.T) {
	assert := assertWrap(t)
	{
		base := FillSliceTest{}
		assert("pass", Fill(&base), nil)
		assert("pass", base.SliceA, []int{})
		assert("pass", base.SliceB, []map[string]string{})
		assert("pass", base.SliceC, []interface{}{})
	}
}
