package shm

import (
	"github.com/haormj/cyber/log"

	"github.com/haormj/cyber/common"
)

type Notifier interface {
	Shutdown()
	Notify(info *ReadableInfo) bool
	Listen(timeoutMs int, info *ReadableInfo) bool
	Type() string
}

func NewNotifier() Notifier {
	notifierType := ConditionNotifierInstance.Type()
	config := common.GlobalDataInstance.Config()
	if config != nil && config.TransportConf != nil &&
		config.TransportConf.ShmConf != nil &&
		config.TransportConf.ShmConf.NotifierType != nil {
		notifierType = config.TransportConf.ShmConf.GetNotifierType()
	}

	log.Logger.Debug("notifier type", "notifierType", notifierType)

	switch notifierType {
	case ConditionNotifierInstance.Type():
		return ConditionNotifierInstance
	case MulticastNotifierInstance.Type():
		return MulticastNotifierInstance
	default:
		log.Logger.Debug("unknown notifier, we use default notifier", "notifierType", notifierType)
		return ConditionNotifierInstance
	}
}
