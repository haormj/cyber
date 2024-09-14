## cyber

golang version of baidu apollo cyber, this is also under development, not for production.

### feature

1. DONE
	- [x] support timer component
	- [x] support data driven component 0-4
	- [x] support run component(only write by go) by `cyber mainboard -d ./xxx.dag` by golang plugin
	- [x] support service by shm
	- [x] support communicate with c++ cyber by shm

2. TODO
	- [ ] support rtps, but https://github.com/eProsima/Fast-DDS/ not have golang version. maybe use cgo
	- [ ] support service discovery
	- [ ] ...

### how to use

#### quick start

1. git clone git@github.com:haormj/cyber.git
2. make
3. cd examples/component/timer/ && make
4. cd examples/component/component4 && make
5. ./output/mainboard -d ./examples/component/timer/dag/timer4.dag -d ./examples/component/component4/dag/component4.dag

#### others

look at ./examples

### reference

- https://github.com/ApolloAuto/apollo