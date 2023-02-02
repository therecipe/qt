package interop

import (
	"encoding/json"
	"sync"

	"github.com/bluszcz/cutego/core"
)

var (
	asyncCallChan = make(chan string, 0)
	syncCallChan  = make(chan string, 0)

	remoteMainThreadIsBlocked      bool
	remoteMainThreadIsBlockedMutex = new(sync.Mutex)
	localMainThreadIsBlocked       bool

	AsyncCallIntoRemote func(string)
	SyncCallIntoRemote  func(string)

	mainThreadHelperChan        = make(chan func(), 0)
	syncCallIntoLocalReturnChan = make(chan string, 0)
)

//
//
//

var MainThreadHelper = NewMainThreadHelper(nil)

type mainThreadHelper struct {
	core.QObject
	_ func(f func()) `slot:"runOnMainThread,auto"`
}

func (h *mainThreadHelper) runOnMainThread(f func()) { f() }

//
//
//

func _syncCallIntoLocal(s string) string {
	defer func() {
		remoteMainThreadIsBlockedMutex.Lock()
		remoteMainThreadIsBlocked = false
		remoteMainThreadIsBlockedMutex.Unlock()
	}()
	remoteMainThreadIsBlockedMutex.Lock()
	remoteMainThreadIsBlocked = true
	remoteMainThreadIsBlockedMutex.Unlock()

	pseudoIn := genPseudoIn(s)

	if pseudoIn.IsArray() && pseudoIn.Property2(0).ToString() == "___return" {
		syncCallIntoRemoteReturnChan <- pseudoIn.Property2(1).ToString()
		return <-syncCallIntoLocalReturnChan
	}

	var ret *PseudoQJSValue
	done := make(chan struct{}, 0)
	if localMainThreadIsBlocked {
		mainThreadHelperChan <- func() { ret = Z_wrapperFunction(pseudoIn); close(done) }
	} else {
		MainThreadHelper.RunOnMainThread(func() { ret = Z_wrapperFunction(pseudoIn); close(done) })
	}
	for {
		select {
		case s := <-syncCallIntoRemoteChan:
			go func() {
				<-done //needed to empty the chan
				syncCallIntoLocalReturnChan <- genPseudoOut(<-Z_wrapperFunctionReturnChan)
			}()

			data, _ := json.Marshal([]string{"___block", s})
			return string(data)

		case <-done:
			goto br
		}
	}
br:

	return genPseudoOut(ret)
}

func _asyncCallIntoRemoteResponse(s string) { asyncCallChan <- s }
func _syncCallIntoRemoteResponse(s string)  { syncCallChan <- s }
