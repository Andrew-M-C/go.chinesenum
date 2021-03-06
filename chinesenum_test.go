package chinesenum

import "testing"

func TestChineseNum(t *testing.T) {
	hans := Get(ZhHans)
	opt := Option{UseOralTwo: true}

	t.Log(hans.Itoa(0))
	t.Log(hans.Itoa(1))
	t.Log(hans.Itoa(12))
	t.Log(hans.Itoa(123))
	t.Log(hans.Itoa(1234))
	t.Log(hans.Itoa(12345))
	t.Log(hans.Itoa(123456))
	t.Log(hans.Itoa(1234567))
	t.Log(hans.Itoa(12345678))
	t.Log(hans.Itoa(123456781))
	t.Log(hans.Itoa(1234567812))
	t.Log(hans.Itoa(12345678123))
	t.Log(hans.Itoa(123456781234))
	t.Log(hans.Itoa(1234567812345))
	t.Log(hans.Itoa(12345678123456))
	t.Log(hans.Itoa(123456781234567))
	t.Log(hans.Itoa(1234567812345678))
	t.Log(hans.Itoa(1))
	t.Log(hans.Itoa(10))
	t.Log(hans.Itoa(100))
	t.Log(hans.Itoa(1000))
	t.Log(hans.Itoa(10000))
	t.Log(hans.Itoa(100000))
	t.Log(hans.Itoa(1000000))
	t.Log(hans.Itoa(10000000))
	t.Log(hans.Itoa(100000000))
	t.Log(hans.Itoa(1000000000))
	t.Log(hans.Itoa(10000000000))
	t.Log(hans.Itoa(100000000000))
	t.Log(hans.Itoa(1000000000000))
	t.Log(hans.Itoa(10000000000000))
	t.Log(hans.Itoa(100000000000000))
	t.Log(hans.Itoa(1000000000000000))
	t.Log(hans.Itoa(2, opt))
	t.Log(hans.Itoa(22, opt))
	t.Log(hans.Itoa(222, opt))
	t.Log(hans.Itoa(2222, opt))
	t.Log(hans.Itoa(22222, opt))
	t.Log(hans.Itoa(222222, opt))
	t.Log(hans.Itoa(2222222, opt))
	t.Log(hans.Itoa(22222222, opt))
	t.Log(hans.Itoa(222222222, opt))
	t.Log(hans.Itoa(2222222222, opt))
	t.Log(hans.Itoa(22222222222, opt))
	t.Log(hans.Itoa(222222222222, opt))
	t.Log(hans.Itoa(2222222222222, opt))
	t.Log(hans.Itoa(22222222222222, opt))
	t.Log(hans.Itoa(222222222222222, opt))
	t.Log(hans.Itoa(2222222222222222, opt))
	t.Log(hans.Itoa(9999999999999999))
	t.Log(hans.Itoa(12000000))
	t.Log(hans.Itoa(12000001))
	t.Log(hans.Itoa(12000010))
	t.Log(hans.Itoa(12000011))
	t.Log(hans.Itoa(12000100))
	t.Log(hans.Itoa(12001000))
	t.Log(hans.Itoa(12010000))
	t.Log(hans.Itoa(12100000))
	t.Log(hans.Itoa(1200000))
}
