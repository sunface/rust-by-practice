package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/mafanr/g"
	"go.uber.org/zap"

	"github.com/mafanr/vgo/util"
)

// SyncCall ...
type SyncCall struct {
	sync.RWMutex
	Chans map[uint32]chan *util.VgoPacket
}

// NewSyncCall ...
func NewSyncCall() *SyncCall {
	return &SyncCall{
		Chans: make(map[uint32]chan *util.VgoPacket),
	}
}

func (sc *SyncCall) addChan(id uint32, packetC chan *util.VgoPacket) {
	sc.Lock()
	defer sc.Unlock()
	sc.Chans[id] = packetC
}

// syncRead 阻塞等待
func (sc *SyncCall) syncRead(id uint32, timeOut int, isStop bool) (*util.VgoPacket, error) {
	sc.RLock()
	packetC, ok := sc.Chans[id]
	sc.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unfind chan, id is %d", id)
	}

	ticker := time.NewTicker(time.Duration(timeOut) * time.Second)
	defer func() {
		ticker.Stop()
		if isStop {
			sc.stopChan(id)
		}
	}()
	select {
	case <-ticker.C:
		g.L.Warn("syncRead:ticker.C.timeout", zap.Uint32("id", id), zap.Int("timeOut", timeOut))
		break
	case packet, ok := <-packetC:
		if ok {
			return packet, nil
		}
		break
	}
	return nil, nil
}

// syncWrite 阻塞写
func (sc *SyncCall) syncWrite(id uint32, packet *util.VgoPacket) error {
	sc.RLock()
	packetC, ok := sc.Chans[id]
	sc.RUnlock()
	if !ok {
		return fmt.Errorf("unfind chan, id is %d", id)
	}
	packetC <- packet
	return nil
}

func (sc *SyncCall) stopChan(id uint32) {
	sc.Lock()
	defer sc.Unlock()
	packetC, ok := sc.Chans[id]
	if ok {
		delete(sc.Chans, id)
		close(packetC)
	}
}
