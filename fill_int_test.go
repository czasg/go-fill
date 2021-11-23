package fill

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

type TestIntFillEnv struct {
	A int   `env:",default=1"`
	B int8  `env:",default=1"`
	C int16 `env:",default=1"`
	D int32 `env:",default=1"`
	E int64 `env:",default=1"`
	F int   `env:",require"`
}

type TestIntFillDefault struct {
	A int   `default:"1"`
	B int8  `default:"1"`
	C int16 `default:"1"`
	D int32 `default:"1"`
	E int64 `default:"xxx"`
}

type TestIntFillEmpty struct {
	A int
	B int8
	C int16
	D int32
	E int64
}

type TestIntFillNotEmpty struct {
	A int
	B int8
	C int16
	D int32
	E int64
}

func TestIntFill(t *testing.T) {
	assert := assertWrap(t)
	nowUnix := time.Now().Unix()
	_ = os.Setenv("A", fmt.Sprintf("%d", nowUnix))
	{
		test := TestIntFillEnv{}
		err := Fill(&test, OptEnv)
		assert("test.env", test.A, int(nowUnix))
		assert("test.env", test.B, int8(1))
		assert("test.env", test.C, int16(1))
		assert("test.env", test.D, int32(1))
		assert("test.env", test.E, int64(1))
		assert("test.err.require", err, errors.New("F require"))
	}
	{
		test := TestIntFillDefault{}
		err := Fill(&test, OptDefault)
		assert("test.default", test.A, 1)
		assert("test.default", test.B, int8(1))
		assert("test.default", test.C, int16(1))
		assert("test.default", test.D, int32(1))
		assert("test.default", test.E, int64(0))
		assert("test.err.invalid", err, errors.New("E invalid [xxx]"))
	}
	{
		test := TestIntFillEmpty{}
		err := Fill(&test)
		assert("test.empty", test.A, 0)
		assert("test.empty", test.B, int8(0))
		assert("test.empty", test.C, int16(0))
		assert("test.empty", test.D, int32(0))
		assert("test.empty", test.E, int64(0))
		assert("test.err.nil", err, nil)
	}
	{
		test := TestIntFillNotEmpty{
			A: 1,
			B: 1,
			C: 1,
			D: 1,
			E: 1,
		}
		err := Fill(&test)
		assert("test.not.empty", test.A, 1)
		assert("test.not.empty", test.B, int8(1))
		assert("test.not.empty", test.C, int16(1))
		assert("test.not.empty", test.D, int32(1))
		assert("test.not.empty", test.E, int64(1))
		assert("test.err.nil", err, nil)
	}
}
