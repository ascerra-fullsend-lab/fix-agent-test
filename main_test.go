package main

import "testing"

func TestGreet(t *testing.T) {
	got := greet("world")
	want := "hello world"
	if got != want {
		t.Errorf("greet(\"world\") = %q, want %q", got, want)
	}
}

func TestGreetEmpty(t *testing.T) {
	got := greet("")
	want := defaultGreeting
	if got != want {
		t.Errorf("greet(\"\") = %q, want %q", got, want)
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{4, 5, 20},
		{0, 10, 0},
	}
	for _, tt := range tests {
		got := multiply(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("multiply(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
