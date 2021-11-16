# Excercism - Go track

These are my personal notes regarding the "Go track" from Excercism.

My initial idea is to go through the multiple exercises and document my learning process here. 

On top of the excercises provided by the platform, I will try to perform the testing for those. Let's see how it goes...


## About Testing

After completeling Gopher's Gorgeous Lasagna excercise(
https://exercism.org/tracks/go/exercises/lasagna ), I decided to give it a try to the testing part since only in theory I was familiar with the concept. This was my initial attempt:

```go
package lasagna

import "testing"

type lasagnaTests struct {
	name     string
	layers   int
	time     int
	expected int
}

func TestOvenTime(t *testing.T) {

	tt := lasagnaTests{
		name:     "Calculates how many minutes the lasagna should be in the oven",
		layers:   0,
		time:     40,
		expected: 40,
	}

	if got := OvenTime; got != tt.expected {
		t.Errorf("OvenTime(%d) = %d; want %d", tt.expected, got, tt.expected)
	}

}
```

Then, I tried like this and got a "PASS":

```
λ go test
PASS
ok      github.com/rabocse/go-ascension/excercism/concept/lasagna       0.135s
```

Before continuing doing more, I wanted to see how Excercism's test might look and found the following:

```go
package lasagna

import "testing"

type lasagnaTests struct {
	name                   string
	layers, time, expected int
}

func TestOvenTime(t *testing.T) {
	tests := []lasagnaTests{
		{
			name:     "Calculates how many minutes the lasagna should be in the oven",
			layers:   0,
			time:     40,
			expected: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OvenTime; got != tt.expected {
				t.Errorf("OvenTime(%d)  = %d; want %d", tt.expected, got, tt.expected)
			}
		})
	}
}
```

The above test used by Excercism raised multiple questions around testing for me, so I will go through the "testing documentation" to clarify those. For example: Why t.Run is used? 

-------------------------------------------------------------

### t.Run
---
I found t.Run in the Go documentation under the section listed as "Subtests and Subenchmarks"

https://pkg.go.dev/testing#hdr-Subtests_and_Sub_benchmarks

- It allows defining subtests and sub-benchmarks, without having to define separate functions for each.
- This enables uses like table-driven benchmarks and creating hierarchical tests. 
- It also provides a way to share common setup and tear-down code. For example:

```go
func TestFoo(t *testing.T) {
    // <setup code>
    t.Run("A=1", func(t *testing.T) { ... })
    t.Run("A=2", func(t *testing.T) { ... })
    t.Run("B=1", func(t *testing.T) { ... })
    // <tear-down code>
}
```

- ___Table-driven tests:___ Instead of mantaining separate test functions, we can build a "table" of input and the corresponding output we expect, then use a single test function to check each entry in the table.
	There is no standard format for the table, ut one common solution is to define a new type, specifically for use in your tests, that holds the input and expected outut for each test.  (Definition taken from "Head First Go" book)

- Notice how each t.Run can run a specific test.


... to be continued...







