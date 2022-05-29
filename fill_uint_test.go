package fill

import (
	"errors"
	"os"
	"testing"
)

type FillUintTest struct {
	UintA uint
	UintB uint  `fill:",default=1"`
	UintC uint  `fill:"UintC"`
	UintD uint  `fill:"UintD,require"`
	UintE uint  `fill:"UintE,empty"`
	UintF *uint `fill:"UintF"`
}

func TestUintFill(t *testing.T) {
	assert := assertWrap(t)
	{
		base := FillUintTest{UintD: 1}
		assert("ignore", Fill(&base), nil)
		assert("ignore", base.UintD, uint(1))
	}
	{
		base := FillUintTest{UintD: 1}
		assert("default", Fill(&base), nil)
		assert("default", base.UintB, uint(1))
	}
	{
		_ = os.Setenv("UintC", "1")
		base := FillUintTest{UintD: 1}
		assert("env", FillEnv(&base), nil)
		assert("env", base.UintC, uint(1))
		_ = os.Unsetenv("UintC")
	}
	{
		base := FillUintTest{}
		assert("require", Fill(&base), errors.New("UintD require"))
	}
	{
		_ = os.Setenv("UintE", "1")
		base := FillUintTest{UintD: 1}
		assert("empty", FillEnv(&base), nil)
		assert("empty", base.UintE, uint(0))
		_ = os.Unsetenv("UintE")
	}
	{
		var v *uint
		base := FillUintTest{UintD: 1}
		assert("ptr", FillEnv(&base), nil)
		assert("ptr", base.UintF, v)
	}
    {
        _ = os.Setenv("UintC", "x")
        base := FillUintTest{UintD: 1}
        assert("env error", FillEnv(&base), errors.New("UintC invalid [x]"))
        _ = os.Unsetenv("UintC")
    }
}
