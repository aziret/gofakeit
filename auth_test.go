package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleUsername() {
	Seed(11)
	fmt.Println(Username())

	// Output: Treutel8125
}

func ExampleFaker_Username() {
	f := New(11)
	fmt.Println(f.Username())

	// Output: Treutel8125
}

func BenchmarkUsername(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Username()
	}
}

func TestPassword(t *testing.T) {
	length := 10

	pass := Password(&PassConfig{Type: Simple, Length: length, Lower: true, Upper: true, Numeric: true, Special: true, Space: true})

	if len(pass) != length {
		t.Error("Password length does not equal requested length")
	}

	// Test fully empty
	pass = Password(&PassConfig{Type: Simple, Length: 0, Lower: false, Upper: false, Numeric: false, Special: false, Space: false})
	if pass == "" {
		t.Error("Password should not be empty")
	}

	// Test it doesn't start or end with a space
	for i := 0; i < 1000; i++ {
		pass = Password(&PassConfig{Type: Simple, Length: 10, Lower: true, Upper: true, Numeric: true, Special: true, Space: true})
		if pass[0] == ' ' || pass[len(pass)-1] == ' ' {
			t.Error("Password should not start or end with a space")
		}
	}
}

func ExamplePassword() {
	Seed(11)
	fmt.Println(Password(&PassConfig{Type: Simple, Length: 32, Lower: true, Upper: false, Numeric: false, Special: false, Space: false}))
	fmt.Println(Password(&PassConfig{Type: Simple, Length: 32, Lower: false, Upper: true, Numeric: false, Special: false, Space: false}))
	fmt.Println(Password(&PassConfig{Type: Simple, Length: 32, Lower: false, Upper: false, Numeric: true, Special: false, Space: false}))
	fmt.Println(Password(&PassConfig{Type: Simple, Length: 32, Lower: false, Upper: false, Numeric: false, Special: true, Space: false}))
	fmt.Println(Password(&PassConfig{Type: Simple, Length: 32, Lower: true, Upper: true, Numeric: true, Special: true, Space: true}))
	fmt.Println(Password(&PassConfig{Type: Simple, Length: 4, Lower: true, Upper: true, Numeric: true, Special: true, Space: true}))

	// Output: cfelntbponnbbzrhswobuwlxajeeclrx
	// KYEKNGUUNKUYSFBUFFTGDKUVCVYKPONP
	// 43622637275953627791234759581343
	// @.__-._-!-!_..!-_*_*__-@*.__.__!
	// -DTHJ@.oF@d@L5F65 N-.@U5xWX F0DI
	// foZnB
}

func ExampleFaker_Password() {
	f := New(11)
	fmt.Println(f.Password(&PassConfig{Type: Simple, Length: 32, Lower: true, Upper: false, Numeric: false, Special: false, Space: false}))
	fmt.Println(f.Password(&PassConfig{Type: Simple, Length: 32, Lower: false, Upper: true, Numeric: false, Special: false, Space: false}))
	fmt.Println(f.Password(&PassConfig{Type: Simple, Length: 32, Lower: false, Upper: false, Numeric: true, Special: false, Space: false}))
	fmt.Println(f.Password(&PassConfig{Type: Simple, Length: 32, Lower: false, Upper: false, Numeric: false, Special: true, Space: false}))
	fmt.Println(f.Password(&PassConfig{Type: Simple, Length: 32, Lower: true, Upper: true, Numeric: true, Special: true, Space: true}))
	fmt.Println(f.Password(&PassConfig{Type: Simple, Length: 4, Lower: true, Upper: true, Numeric: true, Special: true, Space: true}))

	// Output: cfelntbponnbbzrhswobuwlxajeeclrx
	// KYEKNGUUNKUYSFBUFFTGDKUVCVYKPONP
	// 43622637275953627791234759581343
	// @.__-._-!-!_..!-_*_*__-@*.__.__!
	// -DTHJ@.oF@d@L5F65 N-.@U5xWX F0DI
	// foZnB
}

func BenchmarkPassword(b *testing.B) {
	config := PassConfig{
		Type:    Simple,
		Length:  50,
		Lower:   true,
		Upper:   true,
		Numeric: true,
		Special: true,
		Space:   true,
	}
	for i := 0; i < b.N; i++ {
		Password(&config)
	}
}
