package main

import "testing"

func TestIsPrime(t *testing.T) {

	t.Helper()

	f := func(num int, expRes bool) {
		res := IsPrime(num)

		if res != expRes {
			t.Fatalf("unexpected result for IsPrime(%v); got: %t; want: %t", num, res, expRes)
		}
	}

	f(1, false)
	f(2, true)
	f(3, true)
	f(4, false)
	f(5, true)
	f(6, false)

}
