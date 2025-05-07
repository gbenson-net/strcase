package strcase

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSnakeCaseCPUMHz(t *testing.T) {
	s := "cpu MHz"

	for _, tc := range []struct {
		Func   func(string) string
		Expect string
	}{
		{ToSnake, "cpu_mhz"},
		{ToScreamingSnake, "CPU_MHZ"},
		{ToKebab, "cpu-mhz"},
		{ToScreamingKebab, "CPU-MHZ"},
		{ToCamel, "CpuMhz"},
		{ToLowerCamel, "cpuMhz"},
	} {
		name := runtime.FuncForPC(reflect.ValueOf(tc.Func).Pointer()).Name()
		dot := strings.LastIndex(name, ".")
		assert.Assert(t, dot >= 0)
		name = name[dot+1:]

		actual := tc.Func(s)
		msg := fmt.Sprintf("%s: %q", name, actual)
		assert.Check(t, actual == tc.Expect, msg)
	}
}

func TestSnakeCaseMemTotal(t *testing.T) {
	s := "MemTotal"

	for _, tc := range []struct {
		Func   func(string) string
		Expect string
	}{
		{ToSnake, "mem_total"},
		{ToScreamingSnake, "MEM_TOTAL"},
		{ToKebab, "mem-total"},
		{ToScreamingKebab, "MEM-TOTAL"},
		{ToCamel, "MemTotal"},
		{ToLowerCamel, "memTotal"},
	} {
		name := runtime.FuncForPC(reflect.ValueOf(tc.Func).Pointer()).Name()
		dot := strings.LastIndex(name, ".")
		assert.Assert(t, dot >= 0)
		name = name[dot+1:]

		actual := tc.Func(s)
		msg := fmt.Sprintf("%s: %q", name, actual)
		assert.Check(t, actual == tc.Expect, msg)
	}
}
