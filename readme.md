# Always remember the word BET
B - Benchmark
E - Example
T - Test

```
# Functions

func BenchmarkCat(b *testing.B) {}
func ExampleCat() {}
func TestCat(t *testing.T) {}

```

# Commands
```
# Testing
go test

# Benchmarking
go test -bench

# Coverage
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out

```