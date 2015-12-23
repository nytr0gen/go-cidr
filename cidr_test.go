package cidr

import (
    "testing"
)

func TestRange(t *testing.T) {
    r, err := NewRange("127.0.0.0/30")
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "127.0.0.0",
        "127.0.0.1",
        "127.0.0.2",
        "127.0.0.3",
    }

    i := 0
    if r.String() != ips[i] {
        t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
    }

    for r.Next() {
        i++
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }
    }
}

func TestRangeWith31Prefix(t *testing.T) {
    r, err := NewRangeWithBlockSize("127.0.0.0/30", 31)
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "127.0.0.0",
        "127.0.0.2",
    }

    i := 0
    if r.String() != ips[i] {
        t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
    }

    for r.Next() {
        i++
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }
    }
}

func TestRangeWith25Prefix(t *testing.T) {
    r, err := NewRangeWithBlockSize("127.0.0.0/23", 25)
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "127.0.0.0",
        "127.0.0.128",
        "127.0.1.0",
        "127.0.1.128",
    }

    i := 0
    if r.String() != ips[i] {
        t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
    }

    for r.Next() {
        i++
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }
    }
}


func TestRangeWith24Prefix(t *testing.T) {
    r, err := NewRangeWithBlockSize("127.0.0.0/22", 24)
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "127.0.0.0",
        "127.0.1.0",
        "127.0.2.0",
        "127.0.3.0",
    }

    i := 0
    if r.String() != ips[i] {
        t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
    }

    for r.Next() {
        i++
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }
    }
}

func TestRangeWith24PrefixShowingPrefix(t *testing.T) {
    r, err := NewRangeWithBlockSize("127.0.0.0/22", 24)
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "127.0.0.0/24",
        "127.0.1.0/24",
        "127.0.2.0/24",
        "127.0.3.0/24",
    }

    i := 0
    if r.StringPrefix() != ips[i] {
        t.Fatalf("Failed at '%s' != '%s'\n", r.StringPrefix(), ips[i])
    }

    for r.Next() {
        i++
        if r.StringPrefix() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.StringPrefix(), ips[i])
        }
    }
}

func TestShouldFailOnBadCidr(t *testing.T) {
    ipranges := []string{
        "127.0.0.1/31",
        "127.0.1.2/30",
        "127.0.2.127/25",
        "127.0.3.129/25",
    }

    for _, iprange := range ipranges {
        _, err := NewRange(iprange)
        if err != nil && err.Error() == "Invalid cidr" {
            // pass
        } else {
            t.Fatal("didn't fail on %s", iprange)
        }
    }
}

func TestShouldFailOnBadBlockStep(t *testing.T) {
    ipranges := []string{
        "127.0.0.0/31",
        "127.0.1.0/30",
        "127.0.2.0/25",
        "127.0.3.0/25",
    }

    for _, iprange := range ipranges {
        _, err := NewRangeWithBlockSize(iprange, 24)
        if err != nil && err.Error() == "Invalid block size" {
            // pass
        } else {
            t.Fatalf("didn't fail on %s", iprange)
        }
    }
}
