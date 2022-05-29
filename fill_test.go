package fill

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func assertWrap(t *testing.T) func(name string, a, b interface{}) {
	nw := nameWrap()
	return func(name string, a, b interface{}) {
		name = nw(name)
		if !reflect.DeepEqual(a, b) {
			t.Errorf("%s failure! [%v] != [%v]", name, a, b)
		} else {
			t.Logf("%s - pass", name)
		}
	}
}

func nameWrap() func(name string) string {
	count := 0
	lastName := ""
	return func(name string) string {
		if name == lastName {
			count++
		} else {
			count = 0
			lastName = name
		}
		name = strings.ReplaceAll(name, " ", "_")
		return fmt.Sprintf("%s-%d", name, count)
	}
}

func TestFill(t *testing.T) {
	assert := assertWrap(t)
	{
		assert("NotPointerStructErr", Fill(""), NotPointerStructErr)
		assert("NotPointerStructErr", FillEnv(""), NotPointerStructErr)
	}
}
