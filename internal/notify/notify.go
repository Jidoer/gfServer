package notify

import (
	// "errors"
	"gfAdmin/internal/model"
	"sync"
	// "time"
	// "your_project/model"
)

var (
    waiters   = make(map[string]chan *model.AllocatedReq)
    waitersMu sync.Mutex
)

// 注册一个等待
func RegisterWait(uuid string) chan *model.AllocatedReq {
    ch := make(chan *model.AllocatedReq, 1)
    waitersMu.Lock()
    waiters[uuid] = ch
    waitersMu.Unlock()
    return ch
}

// 通知某个 uuid
func Notify(uuid string, data *model.AllocatedReq) {
    waitersMu.Lock()
    ch, ok := waiters[uuid]
    if ok {
        delete(waiters, uuid)
    }
    waitersMu.Unlock()

    if ok {
        ch <- data
        close(ch)
    }
}

// 取消等待（超时或出错）
func Cancel(uuid string) {
    waitersMu.Lock()
    ch, ok := waiters[uuid]
    if ok {
        delete(waiters, uuid)
    }
    waitersMu.Unlock()
    if ok {
        close(ch)
    }
}