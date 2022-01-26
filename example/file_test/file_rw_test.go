package file_test

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// 文件写入测试
func TestFileWrite(t *testing.T) {
	content := []byte("测试1\n测试2\n")
	err := ioutil.WriteFile("./log.txt", content, 0644)
	if err != nil {
		panic(err)
	}
}

// 读取文件内容到内存
func TestFileRead2Mem(t *testing.T) {
	file, err := os.Open("./log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	t.Log(string(content))
}

// 按字节读取
func TestFileReadByByte(t *testing.T) {
	filepath := "./log.txt"
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 0)
	buf := make([]byte, 1024) // 一次读取多少个字节
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		t.Log(string(buf[:n]))
		break
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	t.Log(string(chunks))
}
