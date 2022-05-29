package fill

import (
	"errors"
	"os"
	"testing"
)

type FillFloatTest struct {
	FloatA float32
	FloatB float32  `fill:",default=1"`
	FloatC float32  `fill:"FloatC"`
	FloatD float32  `fill:"FloatD,require"`
	FloatE float32  `fill:"FloatE,empty"`
	FloatF *float32 `fill:"FloatF"`
}

func TestFloatFill(t *testing.T) {
	assert := assertWrap(t)
	{
		base := FillFloatTest{FloatD: 1}
		assert("ignore", Fill(&base), nil)
		assert("ignore", base.FloatD, float32(1))
	}
	{
		base := FillFloatTest{FloatD: 1}
		assert("default", Fill(&base), nil)
		assert("default", base.FloatB, float32(1))
	}
	{
		_ = os.Setenv("FloatC", "1")
		base := FillFloatTest{FloatD: 1}
		assert("env", FillEnv(&base), nil)
		assert("env", base.FloatC, float32(1))
		_ = os.Unsetenv("FloatC")
	}
	{
		base := FillFloatTest{}
		assert("require", Fill(&base), errors.New("FloatD require"))
	}
	{
		_ = os.Setenv("FloatE", "1")
		base := FillFloatTest{FloatD: 1}
		assert("empty", FillEnv(&base), nil)
		assert("empty", base.FloatE, float32(0))
		_ = os.Unsetenv("FloatE")
	}
	{
		var v *float32
		base := FillFloatTest{FloatD: 1}
		assert("ptr", FillEnv(&base), nil)
		assert("ptr", base.FloatF, v)
	}
	{
		_ = os.Setenv("FloatC", "x")
		base := FillFloatTest{FloatD: 1}
		assert("env", FillEnv(&base), errors.New("FloatC invalid [x]"))
		_ = os.Unsetenv("FloatC")
	}
}
