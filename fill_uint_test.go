package fill

import (
    "errors"
    "os"
    "testing"
)

type TestUintFillEnv struct {
    A uint   `env:",default=1"`
    B uint8  `env:",default=1"`
    C uint16 `env:",default=1"`
    D uint32 `env:",default=1"`
    E uint64 `env:",default=1"`
    F uint    `env:",require"`
}

type TestUintFillDefault struct {
    A uint   `default:"1"`
    B uint8  `default:"1"`
    C uint16 `default:"1"`
    D uint32 `default:"1"`
    E uint64 `default:"xxx"`
}

type TestUintFillEmpty struct {
    A uint
    B uint8
    C uint16
    D uint32
    E uint64
}

type TestUiIntFillNotEmpty struct {
    A uint
    B uint8
    C uint16
    D uint32
    E uint64
}

func TestUintFill(t *testing.T) {
    assert := assertWrap(t)
    _ = os.Setenv("A", "")
    _ = os.Setenv("B", "")
    _ = os.Setenv("C", "")
    _ = os.Setenv("D", "")
    _ = os.Setenv("E", "")
    //_ = os.Setenv("F", "")
    {
        test := TestUintFillEnv{}
        err := Fill(&test, OptEnv)
        assert("test.env", test.A, uint(1))
        assert("test.env", test.B, uint8(1))
        assert("test.env", test.C, uint16(1))
        assert("test.env", test.D, uint32(1))
        assert("test.env", test.E, uint64(1))
        assert("test.err.require", err, errors.New("F require"))
    }
    {
        test := TestUintFillDefault{}
        err := Fill(&test, OptDefault)
        assert("test.env", test.A, uint(1))
        assert("test.env", test.B, uint8(1))
        assert("test.env", test.C, uint16(1))
        assert("test.env", test.D, uint32(1))
        assert("test.env", test.E, uint64(0))
        assert("test.err.invalid", err, errors.New("E invalid [xxx]"))
    }
    {
        test := TestUintFillEmpty{}
        err := Fill(&test)
        assert("test.env", test.A, uint(0))
        assert("test.env", test.B, uint8(0))
        assert("test.env", test.C, uint16(0))
        assert("test.env", test.D, uint32(0))
        assert("test.env", test.E, uint64(0))
        assert("test.err.nil", err, nil)
    }
    {
        test := TestUiIntFillNotEmpty{
            A: 2,
            B: 2,
            C: 2,
            D: 2,
            E: 2,
        }
        err := Fill(&test)
        assert("test.env", test.A, uint(2))
        assert("test.env", test.B, uint8(2))
        assert("test.env", test.C, uint16(2))
        assert("test.env", test.D, uint32(2))
        assert("test.env", test.E, uint64(2))
        assert("test.err.nil", err, nil)
    }
}
