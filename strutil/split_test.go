package strutil

import "testing"

func TestUrlSplit(t *testing.T) {
	parse, err := SplitPath("///accounts/account/users/user///")
	t.Log(parse, err)
}

func BenchmarkUrlSplit(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SplitPath("///accounts/account/users/user///")
	}
}

func TestTraverse(t *testing.T) {
	url := "//////"
	length := len(url)
	for part, next := "", 0; next < length; {
		part, next = Between(url, '/', next)
		t.Log(part, next)
	}
}

func BenchmarkTraverse(t *testing.B) {
	url := Trim("///accounts/account////users/user///", '/')
	length := len(url)
	for i := 0; i < t.N; i++ {
		for next := 0; next < length; {
			_, next = Between(url, '/', next)
		}
	}
}

func TestSplitCount(t *testing.T) {
	t.Log(SplitCount("///accounts/account/users/user///", '/'))
}
