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
        IP: IP,
        ipnet: ipnet,
        stepPrefix: "/" + strconv.Itoa(blockSize),
        step: long2ip(1 << uint(32 - blockSize)),
    }, nil
}

type Range struct {
    IP net.IP
    ipnet *net.IPNet
    stepPrefix string
    step net.IP
}

func (r *Range) Next() bool {
    for j := len(r.IP) - 1; j >= 0; j-- {
        t := r.IP[j] + r.step[j]
        if r.IP[j] > t && j > 0 {
            r.IP[j - 1]++
        }
        r.IP[j] = t
    }

    return r.ipnet.Contains(r.IP)
}

func (r *Range) String() string {
    return r.IP.String()
}

func (r *Range) StringPrefix() string {
    return r.IP.String() + r.stepPrefix
}
