package tests

import (
	"reflect"
	"testing"
)

// assert fails the test if the condition is false.
// assert: Gagal jika kondisi yang diberikan bernilai false (tidak terpenuhi).
func Assert(tb testing.TB, condition bool, msg string, v ...any) {
	tb.Helper()
	if !condition {
		tb.Fatalf(msg, v...)
	}
}

// ok fails the test if an err is not nil.
// ok: Gagal jika error tidak bernilai nil (artinya terjadi kesalahan).
func OK(tb testing.TB, err error) {
	tb.Helper() // print runtime caller err
	if err != nil {
		tb.Fatalf("unexpected error: %s", err.Error())
	}
}

// equals fails the test if exp is not equal to act.
// equals: Gagal jika nilai yang diharapkan (exp) tidak sama dengan nilai aktual (act).
func Equals(tb testing.TB, exp, act any) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("exp: %#v\n\n\tgot: %#v", exp, act)
	}
}
