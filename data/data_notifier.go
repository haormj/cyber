package data

import "sync"

var DataNotifierInstance = &DataNotifier{}

type Notifier struct {
	Callback func()
}

type DataNotifier struct {
	notifiesMap      sync.Map
	notifiesMapMutex sync.Mutex
}

func (n *DataNotifier) AddNotifier(channelID uint64, notifier *Notifier) {
	n.notifiesMapMutex.Lock()
	defer n.notifiesMapMutex.Unlock()

	v, ok := n.notifiesMap.Load(channelID)
	if ok {
		notifies := v.([]*Notifier)
		notifies = append(notifies, notifier)
		n.notifiesMap.Store(channelID, notifies)
	} else {
		n.notifiesMap.Store(channelID, []*Notifier{notifier})
	}
}

func (n *DataNotifier) Notify(channelID uint64) bool {
	v, ok := n.notifiesMap.Load(channelID)
	if !ok {
		return false
	}

	notifies := v.([]*Notifier)
	for _, notifier := range notifies {
		notifier.Callback()
	}
	return true
}
