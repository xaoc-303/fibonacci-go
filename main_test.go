package main

import (
    "testing"
)

type testpair struct {
    value int
    result int
}

var tests = []testpair {
    { 4, 3 },
    { 5, 5 },
    { 6, 8 },
}

func TestF1(t *testing.T) {
    for _, pair := range tests {
        v := f1(pair.value)
        if v != pair.result {
            t.Error(
                "For", pair.value,
                "expected", pair.result,
                "got", v,
            )
        }
    }
}

func TestF2(t *testing.T) {
    for _, pair := range tests {
        v := f2(pair.value)
        if v != pair.result {
            t.Error(
                "For", pair.value,
                "expected", pair.result,
                "got", v,
            )
        }
    }
}
