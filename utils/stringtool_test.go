package utils

import "testing"

func Test_String2MD5(t *testing.T) {
	res := String2MD5("123456")
	t.Log(res)
}