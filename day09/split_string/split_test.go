package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	//	测试组
	testGroup := []testCase{
		{"babcbef", "b", []string{"a", "c", "ef"}},
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"abcef", "bc", []string{"a", "ef"}},
		{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	testGroup2 := map[string]testCase{
		"case1": {"babcbef", "b", []string{"a", "c", "ef"}},
		"case2": {"a:b:c", ":", []string{"a", "b", "c"}},
		"case3": {"abcef", "bc", []string{"a", "ef"}},
		"case4": {"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("want: %#v, got: %#v", tc.want, got)
		}
	}

	//	可以用map来运行子测试组

	for name, tc := range testGroup2 {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

//	BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

//	性能比较测试
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}
func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}
func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}
func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}
func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}
