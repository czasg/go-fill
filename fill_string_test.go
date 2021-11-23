package fill

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

type TestStringFillEnv struct {
	A string `env:"A"`
	B string `env:",default=test"`
	C string `env:",require"`
}

type TestStringFillDefault struct {
	A string `default:"test"`
}

type TestStringFillEmpty struct {
	A string
}

type TestStringFillNotEmpty struct {
	A string
}

func TestStringFill(t *testing.T) {
	assert := assertWrap(t)
	randomA := fmt.Sprintf("%d", time.Now().Unix())
	_ = os.Setenv("A", randomA)
	{
		test := TestStringFillEnv{}
		err := Fill(&test, OptEnv)
		assert("test.env", test.A, randomA)
		assert("test.default", test.B, "test")
		assert("test.err.require", err, errors.New("C require"))
	}
	{
		test := TestStringFillDefault{}
		err := Fill(&test, OptDefault)
		assert("test.default", test.A, "test")
		assert("test.err.nil", err, nil)
	}
	{
		test := TestStringFillEmpty{}
		err := Fill(&test)
		assert("test.empty", test.A, "")
		assert("test.err.nil", err, nil)
	}
	{
		test := TestStringFillNotEmpty{A: "not empty"}
		err := Fill(&test, OptEnv, OptDefault)
		assert("test.not.empty", test.A, "not empty")
		assert("test.err.nil", err, nil)
	}
}
