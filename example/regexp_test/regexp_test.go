package regexp_test

import (
	"regexp"
	"testing"
)

var reg = regexp.MustCompile("Gola([a-z]+)g")

func TestRegExp(t *testing.T) {
	str := "Golang regular expressions example"

	// 用来匹配子字符串
	match, err := regexp.MatchString(`^Golang`, str)
	t.Log("Match: ", match, " Error: ", err)

}

func TestRegExpComplie(t *testing.T) {
	str := "Golang expressions example"
	t.Log(reg.FindString(str))
}
