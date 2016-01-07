package cidr

import (
    "testing"
    "math/rand"
    "net"
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


    for i := 0; i < len(ips); i++ {
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }

        if !r.Next() { break; }
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

    for i := 0; i < len(ips); i++  {
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }

        r.Next()
    }
}

func TestLastIPsRangeWith25Prefix(t *testing.T) {
    r, err := NewRangeWithBlockSize("44.44.0.0/16", 25)
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "44.44.254.0",
        "44.44.254.128",
        "44.44.255.0",
        "44.44.255.128",
    }

    for i := 0; i < (512 - len(ips)); i++ {
        r.Next()
    }

    for i := 0; i < len(ips); i++ {
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }

        r.Next()
    }
}

func TestLooping(t *testing.T) {
    r, err := NewRangeWithBlockSize("44.44.0.0/16", 16)
    if err != nil {
        t.Fatal(err)
    }

    for i := 0; i < (1<<16); i++ {
        r.Next()
    }

    if (r.Next()) {
        t.Fatalf("Looping")
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

    for i := 0; i < len(ips); i++ {
        if r.String() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.String(), ips[i])
        }

        r.Next()
    }
}

func TestRangeWith24PrefixShowingPrefix(t *testing.T) {
    r, err := NewRangeWithBlockSize("44.44.0.0/16", 24)
    if err != nil {
        t.Fatal(err)
    }

    ips := []string{
        "44.44.0.0/24",
        "44.44.1.0/24",
        "44.44.2.0/24",
        "44.44.3.0/24",
    }

    for i := 0; i < len(ips); i++ {
        if r.StringPrefix() != ips[i] {
            t.Fatalf("Failed at %s != %s\n", r.StringPrefix(), ips[i])
        }

        r.Next()
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

func TestIP2Long(t *testing.T) {
    testsNum := 32
    for testsNum > 0 {
        ip := net.IPv4(byte(rand.Intn(256)),
            byte(rand.Intn(256)),
            byte(rand.Intn(256)),
            byte(rand.Intn(256)))
        assertIP := Long2IP(IP2Long(ip))
        if !ip.Equal(assertIP) {
            t.Fatalf("didn't work %s != %s", ip.String(), assertIP.String())
            return
        }

        testsNum--
    }
}

func TestLong2IP(t *testing.T) {
    testsNum := 32
    for testsNum > 0 {
        long := uint(rand.Uint32())
        assertLong := IP2Long(Long2IP(long))
        if long != assertLong {
            t.Fatalf("didn't work %u != %u", long, assertLong)
            return
        }

        testsNum--
    }
}
