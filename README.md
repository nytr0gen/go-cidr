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

for {
    println(r.String()) // 127.0.0.0
    println(r.StringPrefix()) // 127.0.0.0/31

    if !r.Next() {
        break
    }
}
```

```go
import "github.com/nytr0gen/go-cidr"

ips, err := cidr.List("127.0.0.0/30")
if err != nil {
    panic(err)
}

for _, ip := ips {
    println(ip) // 127.0.0.0
}
```
# go-cidr
