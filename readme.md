# Always remember the word BET
B - Benchmark
E - Example
T - Test

```
BenchmarkCat(b *testing.B)
ExampleCat()
TestCat(t *testing.T)
```

# Commands
```
go test
go test -bench
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```