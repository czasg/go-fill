package fill

import (
	"errors"
	"os"
	"testing"
)

type FillStringTest struct {
	StringA string
	StringB string  `fill:",default=StringB"`
	StringC string  `fill:"StringC"`
	StringD string  `fill:"StringD,require"`
	StringE string  `fill:"StringE,empty"`
	StringF *string `fill:"StringF"`
}

func TestStringFill(t *testing.T) {
	assert := assertWrap(t)
	{
		base := FillStringTest{StringD: "StringD"}
		assert("ignore", Fill(&base), nil)
		assert("ignore", base.StringA, "")
	}
	{
		base := FillStringTest{StringD: "StringD"}
		assert("default", Fill(&base), nil)
		assert("default", base.StringB, "StringB")
	}
	{
		_ = os.Setenv("StringC", "StringC")
		base := FillStringTest{StringD: "StringD"}
		assert("env", FillEnv(&base), nil)
		assert("env", base.StringC, "StringC")
		_ = os.Unsetenv("StringC")
	}
	{
		base := FillStringTest{}
		assert("require", Fill(&base), errors.New("StringD require"))
	}
	{
		_ = os.Setenv("StringE", "StringE")
		base := FillStringTest{StringD: "StringD"}
		assert("empty", FillEnv(&base), nil)
		assert("empty", base.StringE, "")
		_ = os.Unsetenv("StringE")
	}
	{
		var v *string
		base := FillStringTest{StringD: "StringD"}
		assert("ptr", FillEnv(&base), nil)
		assert("ptr", base.StringF, v)
	}
}
