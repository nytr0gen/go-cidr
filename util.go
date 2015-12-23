package cidr

import (
    "net"
)

func init() {
    var privateRanges [3]net.IPNet
    _, privateRanges[0], err := net.ParseCIDR("10.0.0.0/8")
    _, privateRanges[0], err := net.ParseCIDR("172.16.0.0/12")
    _, privateRanges[0], err := net.ParseCIDR("192.168.0.0/16")
    // "10.0.0.0/8"
    // "172.16.0.0/12"
    // "192.168.0.0/16"
}

func long2ip(long uint) net.IP {
    return net.IPv4(byte(long >> 24), byte(long >> 16), byte(long >> 8), byte(long))
}

func IsPrivate(ip net.IP) bool {
    switch {
    case ip[0] == 10: // 10.0.0.0/8: 10.0.0.0 - 10.255.255.255
        return true
    case ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32): // 172.16.0.0/12: 172.16.0.0 - 172.31.255.255
        return true
    case ip[0] == 192 && ip[1] == 168: // 192.168.0.0/16: 192.168.0.0 - 192.168.255.255
        return true
    }

    return false
}
