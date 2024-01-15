//go:build linux

package flags

import (
	"net"
	"unsafe"

	"golang.org/x/sys/unix"
)

// GetAllFlags 获取所有网卡的flags
func GetAllFlags() (map[string][]string, error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	iff := map[uint16]string{
		unix.IFF_UP:          "UP",
		unix.IFF_BROADCAST:   "BROADCAST",
		unix.IFF_DEBUG:       "DEBUG",
		unix.IFF_LOOPBACK:    "LOOPBACK",
		unix.IFF_POINTOPOINT: "POINTOPOINT",
		unix.IFF_NOTRAILERS:  "NOTRAILERS",
		unix.IFF_RUNNING:     "RUNNING",
		unix.IFF_NOARP:       "NOARP",
		unix.IFF_PROMISC:     "PROMISC",
		unix.IFF_ALLMULTI:    "ALLMULIT",
		unix.IFF_MASTER:      "MASTER",
		unix.IFF_SLAVE:       "SLAVE",
		unix.IFF_MULTICAST:   "MULTICAST",
		unix.IFF_PORTSEL:     "PROTSEL",
		unix.IFF_AUTOMEDIA:   "AUTOMEDIA",
		unix.IFF_DYNAMIC:     "DYNAMIC",
	}
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		return nil, err
	}
	defer func(fd int) {
		_ = unix.Close(fd)
	}(fd)
	res := make(map[string][]string)
	for _, val := range ifs {
		var ifReq [40]byte
		copy(ifReq[:], val.Name)
		_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), unix.SIOCGIFFLAGS, uintptr(unsafe.Pointer(&ifReq)))
		if errno != 0 {
			return nil, errno
		}
		flags := *(*uint16)(unsafe.Pointer(&ifReq[16]))
		for flag, status := range iff {
			if flags&flag != 0 {
				res[val.Name] = append(res[val.Name], status)
			}
		}
	}

	return res, nil
}

// GetFlagsByName 通过网卡名获取所有的flags
func GetFlagsByName(name string) ([]string, error) {
	iff := map[uint16]string{
		unix.IFF_UP:          "UP",
		unix.IFF_BROADCAST:   "BROADCAST",
		unix.IFF_DEBUG:       "DEBUG",
		unix.IFF_LOOPBACK:    "LOOPBACK",
		unix.IFF_POINTOPOINT: "POINTOPOINT",
		unix.IFF_NOTRAILERS:  "NOTRAILERS",
		unix.IFF_RUNNING:     "RUNNING",
		unix.IFF_NOARP:       "NOARP",
		unix.IFF_PROMISC:     "PROMISC",
		unix.IFF_ALLMULTI:    "ALLMULIT",
		unix.IFF_MASTER:      "MASTER",
		unix.IFF_SLAVE:       "SLAVE",
		unix.IFF_MULTICAST:   "MULTICAST",
		unix.IFF_PORTSEL:     "PROTSEL",
		unix.IFF_AUTOMEDIA:   "AUTOMEDIA",
		unix.IFF_DYNAMIC:     "DYNAMIC",
	}
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		return nil, err
	}
	defer func(fd int) {
		_ = unix.Close(fd)
	}(fd)
	res := make([]string, 0)

	var ifReq [40]byte
	copy(ifReq[:], name)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), unix.SIOCGIFFLAGS, uintptr(unsafe.Pointer(&ifReq)))
	if errno != 0 {
		return nil, errno
	}
	flags := *(*uint16)(unsafe.Pointer(&ifReq[16]))
	for flag, status := range iff {
		if flags&flag != 0 {
			res = append(res, status)
		}
	}
	return res, nil
}
