package support

import (
	"testing"
)

func TestFirstFirst(t *testing.T) {
    num := FindFirstNumber("thisoneheretwothree4")
    want := 1
    if num != 1 {
        t.Fatalf(`FindFirstNumber("thisonehere") = %q, want match for %#q`, num, want)
    }
    num = FindFirstNumber("nineten")
    want = 9
    if num != 9 {
        t.Fatalf(`FindFirstNumber("thisonehere") = %q, want match for %#q`, num, want)
    }
    num = FindFirstNumber("7onefour1eighttwo5three")
    want = 7
    if num != 7 {
        t.Fatalf(`FindFirstNumber("7onefour1eighttwo5three") = %q, want match for %#q`, num, want)
    }
}

func TestFirstFirst2(t *testing.T) {
    num := FindFirstNumber("thistwohere")
    want := 2
    if num != 2 {
        t.Fatalf(`FindFirstNumber("thistwohere") = %q, want match for %#q`, num, want)
    }
}

func TestFirstFirst3(t *testing.T) {
    num := FindFirstNumber("1thistwohere")
    want := 1
    if num != 1 {
        t.Fatalf(`FindFirstNumber("1thistwohere") = %q, want match for %#q`, num, want)
    }
}



func TestFirstLast(t *testing.T) {
    num := FindLastNumber("1thistwohere")
    want := 2
    if num != 2 {
        t.Fatalf(`FindLastNumber("1thistwohere") = %q, want match for %#q`, num, want)
    }
    num = FindLastNumber("1thistwohere3")
    want = 3
    if num != 3 {
        t.Fatalf(`FindLastNumber("1thistwohere3") = %q, want match for %#q`, num, want)
    }
    num = FindLastNumber("7onefour1eighttwo5three")
    want = 3
    if num != want {
        t.Fatalf(`FindLastNumber("7onefour1eighttwo5three") = %q, want match for %#q`, num, want)
    }
}
