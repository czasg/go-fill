package fill

import (
    "testing"
)

type TestPtrFillNil struct {
    A *string
    B *int
    C map[string]string
    D []string
    E chan struct{}
    F *[2]string
}

type TestPtrFillEnv struct {
    TestPtrFillEnv1 *TestPtrFillEnv1
    TestPtrFillEnv2 *TestPtrFillEnv2
}

type TestPtrFillEnv1 struct {
    A               string `env:",default=test"`
    B               int    `env:",default=1"`
    TestPtrFillEnv2 *TestPtrFillEnv2
}

type TestPtrFillEnv2 struct {
    A string `env:",default=test2"`
    B int    `env:",default=2"`
}

func TestPtrFill(t *testing.T) {
    assert := assertWrap(t)
    {
        test := TestPtrFillNil{}
        err := Fill(&test)
        assert("test fill zero value", *test.A, "")
        assert("test fill zero value", *test.B, 0)
        assert("test fill zero value", test.C, make(map[string]string))
        assert("test fill zero value", test.D, make([]string, 0, 0))
        assert("test fill zero value", len(test.E), 0)
        assert("test fill zero value", err, nil)
    }
    {
        name := "test"
        test := TestPtrFillNil{A: &name}
        err := Fill(&test)
        assert("test fill zero value", *test.A, "test")
        assert("test fill zero value", *test.F, [2]string{})
        assert("test fill zero value", err, nil)
    }
    {
        test := TestPtrFillEnv{}
        err := Fill(&test, OptEnv)
        assert("test fill env", test.TestPtrFillEnv1.A, "")
        assert("test fill env", test.TestPtrFillEnv1.B, 0)
        assert("test fill env", test.TestPtrFillEnv2.A, "")
        assert("test fill env", test.TestPtrFillEnv2.B, 0)
        assert("test fill env", err, nil)
    }
}
