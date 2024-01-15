# flags
> linux系统下获取网卡的flags
- 可以执行ifconfig和执行的结果进行比对
- env
  - os: centos8
  - go version: go1.21.5
- usage: 
  - `go get github.com/kamiertop/flags`
  - `flags.GetAllFlags()`
  - res: "UP", "BROADCAST", "RUNNING", "MULTICAST"