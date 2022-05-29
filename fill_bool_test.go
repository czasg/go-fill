package fill

import (
	"errors"
	"os"
	"testing"
)

type FillBoolTest struct {
	BoolA bool
	BoolB bool  `fill:",default=1"`
	BoolC bool  `fill:"BoolC"`
	BoolD bool  `fill:"BoolD,require"`
	BoolE bool  `fill:"BoolE,empty"`
	BoolF *bool `fill:"BoolF"`
}

func TestBoolFill(t *testing.T) {
	assert := assertWrap(t)
	{
		base := FillBoolTest{BoolD: true}
		assert("ignore", Fill(&base), nil)
		assert("ignore", base.BoolD, true)
	}
	{
		base := FillBoolTest{BoolD: true}
		assert("default", Fill(&base), nil)
		assert("default", base.BoolB, true)
	}
	{
		_ = os.Setenv("BoolC", "1")
		base := FillBoolTest{BoolD: true}
		assert("env", FillEnv(&base), nil)
		assert("env", base.BoolC, true)
		_ = os.Unsetenv("BoolC")
	}
	{
		base := FillBoolTest{}
		assert("require", Fill(&base), errors.New("BoolD require"))
	}
	{
		_ = os.Setenv("BoolE", "true")
		base := FillBoolTest{BoolD: true}
		assert("empty", FillEnv(&base), nil)
		assert("empty", base.BoolE, false)
		_ = os.Unsetenv("BoolE")
	}
	{
		var v *bool
		base := FillBoolTest{BoolD: true}
		assert("ptr", FillEnv(&base), nil)
		assert("ptr", base.BoolF, v)
	}
	{
		_ = os.Setenv("BoolC", "x")
		base := FillBoolTest{BoolD: true}
		assert("env error", FillEnv(&base), errors.New("BoolC invalid [x]"))
		_ = os.Unsetenv("BoolC")
	}
}
