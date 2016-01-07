package cidr

import (
    "net"
    "strconv"
    "strings"
    "errors"
)

func NewRange(iprange string) (*Range, error) {
    return NewRangeWithBlockSize(iprange, 32)
}

func NewRangeWithBlockSize(iprange string, blockSize int) (*Range, error) {
    ip, ipnet, err := net.ParseCIDR(iprange)
    if err != nil {
        return nil, err
    } else if !ip.Equal(ipnet.IP) {
        return nil, errors.New("Invalid cidr")
    }

    var iprangeMask int
    if slashPos := strings.LastIndex(iprange, "/"); slashPos == -1 {
        iprangeMask = 32
    } else {
        iprangeMask, err = strconv.Atoi(iprange[slashPos+1:])
        if err != nil {
            return nil, err
        }
    }

    if err != nil {
        return nil, err
    } else if iprangeMask > blockSize || blockSize > 32 {
        return nil, errors.New("Invalid block size")
    }

    return &Range{
        ip: ip,
        iplong: IP2Long(ip),
        ipnet: ipnet,
        stepprefix: "/" + strconv.Itoa(blockSize),
        lastiplong: IP2Long(ip) + (1 << uint(32 - iprangeMask)) - 1,
        step: 1 << uint(32 - blockSize),
    }, nil
}

type Range struct {
    ipnet *net.IPNet
    ip net.IP
    iplong uint
    lastiplong uint
    stepprefix string
    step uint
}

func (r *Range) Next() bool {
    r.iplong += r.step
    r.ip = Long2IP(r.iplong)

    return r.iplong <= r.lastiplong
}

func (r *Range) String() string {
    return r.ip.String()
}

func (r *Range) StringPrefix() string {
    return r.ip.String() + r.stepprefix
}
