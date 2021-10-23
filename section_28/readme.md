# Golang tests :

- Test
- Example
- Benchmark

```
Function to be tested : Join()
- func TestJoin(t *testing.T){}
- func ExampleJoin(){}
- func BenchmarkJoin(b *testing.B){}
```
# Commands:
```
godoc -http=:8080
go test <dir | filename>
go test -bench <dir | filename>
go test -cover <dir | filename>
go test -coverprofile c.out =>cover profile file.
go tool cover -html=c.out

