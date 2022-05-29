package fill

import (
	"errors"
	"os"
	"testing"
)

type FillIntTest struct {
	IntA int
	IntB int  `fill:",default=1"`
	IntC int  `fill:"IntC"`
	IntD int  `fill:"IntD,require"`
	IntE int  `fill:"IntE,empty"`
	IntF *int `fill:"IntF"`
}

func TestIntFill(t *testing.T) {
	assert := assertWrap(t)
	{
		base := FillIntTest{IntD: 1}
		assert("ignore", Fill(&base), nil)
		assert("ignore", base.IntD, 1)
	}
	{
		base := FillIntTest{IntD: 1}
		assert("default", Fill(&base), nil)
		assert("default", base.IntB, 1)
	}
	{
		_ = os.Setenv("IntC", "1")
		base := FillIntTest{IntD: 1}
		assert("env", FillEnv(&base), nil)
		assert("env", base.IntC, 1)
		_ = os.Unsetenv("IntC")
	}
	{
		base := FillIntTest{}
		assert("require", Fill(&base), errors.New("IntD require"))
	}
	{
		_ = os.Setenv("IntE", "1")
		base := FillIntTest{IntD: 1}
		assert("empty", FillEnv(&base), nil)
		assert("empty", base.IntE, 0)
		_ = os.Unsetenv("IntE")
	}
	{
		var v *int
		base := FillIntTest{IntD: 1}
		assert("ptr", FillEnv(&base), nil)
		assert("ptr", base.IntF, v)
	}
	{
		_ = os.Setenv("IntC", "x")
		base := FillIntTest{IntD: 1}
		assert("env", FillEnv(&base), errors.New("IntC invalid [x]"))
		_ = os.Unsetenv("IntC")
	}
}
