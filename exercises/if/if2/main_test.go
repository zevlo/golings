// if2
// Make me compile!

package main_test

import "testing"

func fooIfFizz(input string) string {
    if input == "fizz" {
        return "foo"
    } else if input == "fuzz" {
        return "bar"
    } else {
        return "baz"
    }
}


func TestFooForFizz(t *testing.T) {
	result := fooIfFizz("fizz")
	if result != "foo" {
		t.Errorf("should be 'foo' but got %s", result)
	}
}

func TestBarForFuzz(t *testing.T) {
	result := fooIfFizz("fuzz")
	if result != "bar" {
		t.Errorf("should be 'bar' but got %s", result)
	}
}

func TestDefaultForBazz(t *testing.T) {
	result := fooIfFizz("random stuff")
	if result != "baz" {
		t.Errorf("should be 'baz' but got %s", result)
	}
}
