package common

import (
	"net"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/pb"
	"github.com/spf13/cast"
)

const kEmptyString = ""

var GlobalDataInstance *GlobalData = NewGlobalData()

type GlobalData struct {
	config        *pb.CyberConfig
	hostIP        string
	hostName      string
	processID     int
	ProcessGroup  string
	ComponentNums int
	SchedName     string
	runMode       pb.RunMode
	clockMode     pb.ClockMode
	// c++ 使用的是 AtomicHashMap，后续可研究替换
	nodeIDMap    sync.Map
	channelIDMap sync.Map
	serviceIDMap sync.Map
	taskIDMap    sync.Map
}

func NewGlobalData() *GlobalData {
	d, err := NewNewGlobalDataE()
	if err != nil {
		panic(err)
	}
	return d
}

func NewNewGlobalDataE() (*GlobalData, error) {
	d := &GlobalData{
		SchedName: "CYBER_DEFAULT",
	}

	if err := d.initHostInfo(); err != nil {
		return nil, err
	}

	// if !d.initConfig() {
	// 	return nil, errors.New("initConfig failed")
	// }

	d.processID = os.Getpid()
	programPath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	if len(programPath) != 0 {
		d.ProcessGroup = GetFileName(programPath) + "_" + cast.ToString(d.processID)
	} else {
		d.ProcessGroup = "cyber_default_" + cast.ToString(d.processID)
	}

	// d.runMode = *d.config.RunModeConf.RunMode
	// d.clockMode = *d.config.RunModeConf.ClockMode

	return d, nil
}

func (d *GlobalData) initHostInfo() error {
	name, err := os.Hostname()
	if err != nil {
		return err
	}
	d.hostName = name
	d.hostIP = "127.0.0.1"

	ip := CyberIP()
	if !strings.HasPrefix(ip, "127") {
		d.hostIP = ip
		return nil
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 ||
			iface.Flags&net.FlagRunning == 0 ||
			iface.Flags&net.FlagLoopback > 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return err
		}

		if len(addrs) == 0 {
			continue
		}

		splits := strings.Split(addrs[0].String(), "/")
		if len(splits) != 2 {
			continue
		}

		if strings.HasPrefix(splits[0], "127") {
			continue
		}

		d.hostIP = splits[0]
		break
	}

	return nil
}

func (d *GlobalData) initConfig() bool {
	p := path.Join(WorkRoot(), "conf/cyber.pb.conf")
	var conf pb.CyberConfig
	if err := GetProtoFromFile(p, &conf); err != nil {
		log.Logger.Error("read cyber default conf failed!", "err", err)
		return false
	}
	d.config = &conf

	return true
}

func (d *GlobalData) ProcessID() int {
	return d.processID
}

func (d *GlobalData) HostIP() string {
	return d.hostIP
}

func (d *GlobalData) HostName() string {
	return d.hostName
}

func (d *GlobalData) Config() *pb.CyberConfig {
	return d.config
}

func (d *GlobalData) EnableSimulationMode() {
	d.runMode = pb.RunMode_MODE_SIMULATION
}

func (d *GlobalData) DisableSimulationMode() {
	d.runMode = pb.RunMode_MODE_REALITY
}

func (d *GlobalData) IsRealityMode() bool {
	return d.runMode == pb.RunMode_MODE_REALITY
}

func (d *GlobalData) IsMockTimeMode() bool {
	return d.clockMode == pb.ClockMode_MODE_MOCK
}

func (d *GlobalData) GenerateHashID(name string) uint64 {
	return Hash([]byte(name))
}

func (d *GlobalData) RegisterNode(nodeName string) uint64 {
	id := Hash([]byte(nodeName))
	for v, ok := d.nodeIDMap.Load(id); ok; id++ {
		if nodeName == v.(string) {
			break
		}

		log.Logger.Warn("Node name hash collision", "nodeName", nodeName, "value", v)
	}
	d.nodeIDMap.Store(id, nodeName)
	return id
}

func (d *GlobalData) GetNodeByID(id uint64) string {
	v, ok := d.nodeIDMap.Load(id)
	if !ok {
		return kEmptyString
	}
	return v.(string)
}

func (d *GlobalData) RegisterChannel(channel string) uint64 {
	id := Hash([]byte(channel))
	for v, ok := d.channelIDMap.Load(id); ok; id++ {
		if channel == v.(string) {
			break
		}

		log.Logger.Warn("Channel name hash collision", "channel", channel, "value", v)
	}
	d.channelIDMap.Store(id, channel)
	return id
}

func (d *GlobalData) GetChannelByID(id uint64) string {
	v, ok := d.channelIDMap.Load(id)
	if !ok {
		return kEmptyString
	}
	return v.(string)
}

func (d *GlobalData) RegisterService(service string) uint64 {
	id := Hash([]byte(service))
	for v, ok := d.serviceIDMap.Load(id); ok; id++ {
		if service == v.(string) {
			break
		}

		log.Logger.Warn("Service name hash collision", "service", service, "value", v)
	}
	d.serviceIDMap.Store(id, service)
	return id
}

func (d *GlobalData) GetServiceByID(id uint64) string {
	v, ok := d.serviceIDMap.Load(id)
	if !ok {
		return kEmptyString
	}
	return v.(string)
}

func (d *GlobalData) RegisterTaskName(taskName string) uint64 {
	id := Hash([]byte(taskName))
	for v, ok := d.taskIDMap.Load(id); ok; id++ {
		if taskName == v.(string) {
			break
		}

		log.Logger.Warn("Service name hash collision", "taskName", taskName, "value", v)
	}
	d.taskIDMap.Store(id, taskName)
	return id
}

func (d *GlobalData) GetTaskNameByID(id uint64) string {
	v, ok := d.taskIDMap.Load(id)
	if !ok {
		return kEmptyString
	}
	return v.(string)
}
