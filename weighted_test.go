package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleWeighted() {
	Seed(11)

	options := []interface{}{"hello", 2, 6.9}
	weights := []float32{1, 2, 3}
	option, _ := Weighted(options, weights)

	fmt.Println(option)
	// Output: 2
}

func TestWeighted(t *testing.T) {
	percOfValue := func(options []interface{}, option interface{}) float32 {
		var count float32 = 0
		for _, o := range options {
			if option == o {
				count++
			}
		}

		return (count / float32(len(options))) * 100
	}

	Seed(11)
	options := []interface{}{"hello", 2, 6.9}
	weights := []float32{1, 2, 3}

	foundOptions := []interface{}{}
	for i := 0; i < 100000; i++ {
		o, _ := Weighted(options, weights)
		foundOptions = append(foundOptions, o)
	}

	perc := percOfValue(foundOptions, "hello")
	if perc < 14 || perc > 18 {
		t.Error("hello was not in the bounds of expected range")
	}
	perc = percOfValue(foundOptions, 2)
	if perc < 30 || perc > 35 {
		t.Error("2 was not in the bounds of expected range")
	}
	perc = percOfValue(foundOptions, 6.9)
	if perc < 48 || perc > 52 {
		t.Error("6.9 was not in the bounds of expected range")
	}
}

func BenchmarkWeighted(b *testing.B) {
	options := []interface{}{"hello", 2, 6.9}
	weights := []float32{1, 2, 3}

	Seed(11)
	for i := 0; i < b.N; i++ {
		Weighted(options, weights)
	}
}
