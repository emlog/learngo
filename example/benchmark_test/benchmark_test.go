// 性能测试
package benchmark_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试：字符串连接
func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("12345", ret)
}

// 测试：使用 buf.WriteString来连接字符串
func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}

// 基准测试：字符串连接
func BenchmarkConcatStringByAdd(b *testing.B) {

	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

// 基准测试：使用 buf.WriteString来连接字符串
func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()

}

// 运行文件里的所有benchmark测试：go test -bench=.
// 运行文件里的所有benchmark测试，并展示内存使用情况：go test -bench=. -benchmem
// B/op:每次操作分配的内存字节数
// allocs/op:每次操作分配内存的次数
