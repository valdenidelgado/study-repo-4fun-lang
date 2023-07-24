package main

import "testing"

func TestGreetings(t *testing.T) {
    result := Greetings()
    expected := "Hello, World!"

    if result != expected {
        t.Errorf("result '%s', expected '%s'", result, expected)
    }
}