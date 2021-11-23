package fill

import (
    "errors"
    "testing"
)

type TestFloatFillEnv struct {
    A float32 `env:",default=1"`
    B float64 `env:",default=1"`
    F float32 `env:",require"`
}

type TestFloatFillDefault struct {
    A float32 `default:"1"`
    B float64 `default:"1"`
    E float32 `default:"xxx"`
}

type TestFloatFillEmpty struct {
    A float32 `default:"1"`
    B float64 `default:"1"`
}

type TestFloatFillNotEmpty struct {
    A float32 `default:"1"`
    B float64 `default:"1"`
}

func TestFloatFill(t *testing.T) {
    assert := assertWrap(t)
    {
        test := TestFloatFillEnv{}
        err := Fill(&test, OptEnv)
        assert("test.env", test.A, float32(1))
        assert("test.env", test.B, float64(1))
        assert("test.err.require", err, errors.New("F require"))
    }
    {
        test := TestFloatFillDefault{}
        err := Fill(&test, OptDefault)
        assert("test.env", test.A, float32(1))
        assert("test.env", test.B, float64(1))
        assert("test.err.require", err, errors.New("E invalid [xxx]"))
    }
    {
        test := TestFloatFillEmpty{}
        err := Fill(&test)
        assert("test.env", test.A, float32(0))
        assert("test.env", test.B, float64(0))
        assert("test.err.require", err, nil)
    }
    {
        test := TestFloatFillNotEmpty{
            A: 1,
            B: 1,
        }
        err := Fill(&test)
        assert("test.env", test.A, float32(1))
        assert("test.env", test.B, float64(1))
        assert("test.err.require", err, nil)
    }
}
