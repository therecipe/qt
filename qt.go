package qt

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	Logger = log.New(os.Stderr, "", log.Ltime)

	signals      = make(map[string]interface{})
	signalsMutex = new(sync.Mutex)

	ids int
)

func init() { runtime.LockOSThread() }

func ExistsSignal(name, signal string) bool {
	signalsMutex.Lock()
	_, exists := signals[fmt.Sprintf("%v:%v", name, signal)]
	signalsMutex.Unlock()
	return exists
}

func LendSignal(name, signal string) interface{} {
	var s interface{}
	signalsMutex.Lock()
	s = signals[fmt.Sprintf("%v:%v", name, signal)]
	signalsMutex.Unlock()
	return s
}

func GetSignal(name, signal string) interface{} {
	if signal == "destroyed" || signal == "deleteLater" || strings.HasPrefix(signal, "~") {
		defer DisconnectAllSignals(name)
	}
	return LendSignal(name, signal)
}

func ConnectSignal(name, signal string, function interface{}) {
	signalsMutex.Lock()
	signals[fmt.Sprintf("%v:%v", name, signal)] = function
	signalsMutex.Unlock()
}

func DisconnectSignal(name, signal string) {
	signalsMutex.Lock()
	delete(signals, fmt.Sprintf("%v:%v", name, signal))
	signalsMutex.Unlock()
}

func DisconnectAllSignals(name string) {
	signalsMutex.Lock()
	for entry := range signals {
		if strings.HasPrefix(entry, fmt.Sprintf("%v:", name)) {
			delete(signals, entry)
		}
	}
	signalsMutex.Unlock()
}

func Identifier() string {
	ids++
	return fmt.Sprint(ids)
}

func DumpSignals() {
	Debug("##############################\tSIGNALSTABLE_START\t##############################")
	signalsMutex.Lock()
	for entry := range signals {
		Debug(entry)
	}
	signalsMutex.Unlock()
	Debug("##############################\tSIGNALSTABLE_END\t##############################")
}

func GoBoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Recover(fn string) {
	if recover() != nil {
		Debug("RECOVERED:", fn)
	}
}

func Debug(fn ...interface{}) {
	if strings.ToLower(os.Getenv("QT_DEBUG")) == "true" {
		Logger.Println(fn...)
	}
}
