## go-cidr

### Examples
```go
import "github.com/nytr0gen/go-cidr"

r, err := cidr.NewRange("127.0.0.0/31")
if err != nil {
    panic(err)
}

println(r.String()) // 127.0.0.0
r.Next() // true
println(r.String()) // 127.0.0.1
r.Next() // false
```

```go
import "github.com/nytr0gen/go-cidr"

r, err := cidr.NewRangeWithBlockSize("127.0.0.0/30", 31)
if err != nil {
    panic(err)
}

println(r.String()) // 127.0.0.0
println(r.StringPrefix()) // 127.0.0.0/31
r.Next() // true

println(r.String()) // 127.0.0.2
println(r.StringPrefix()) // 127.0.0.2/31
r.Next() // false
```
# go-cidr
