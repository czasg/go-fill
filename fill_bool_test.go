package fill

import (
    "errors"
    "testing"
)

type TestFillBoolEnv struct {
    A bool
    B bool `env:",default=true"`
    C bool `env:",default=false"`
    D bool `env:",require"`
}

func TestFillBool(t *testing.T) {
    assert := assertWrap(t)
    {
        test := TestFillBoolEnv{C: true}
        err := Fill(&test, OptEnv)
        assert("test.env", test.A, false)
        assert("test.env", test.B, true)
        assert("test.env", test.C, true)
        assert("test.env", err, errors.New("D require"))
    }
}
