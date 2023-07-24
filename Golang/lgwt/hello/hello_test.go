package main

import "testing"

func TestGreetings(t *testing.T) {
    result := Greetings("Valdeni")
    expected := "Hello, Valdeni!"

    if result != expected {
        t.Errorf("result '%s', expected '%s'", result, expected)
    }
}
